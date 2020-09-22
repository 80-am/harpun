package cmd

import (
	"regexp"
	"strings"

	"github.com/80-am/harpun/db"
	"github.com/gocolly/colly/v2"
)

// Stock to fetch data for
type Stock struct {
	ID     string
	Name   string
	Ticker string
}

// AddStock to db
func AddStock(s Stock) {
	stmt := db.Prepare("INSERT INTO stocks(avanzaId, ticker, name) VALUES(?, ?, ?)")
	stmt.Exec(s.ID, s.Ticker, s.Ticker)
	InfoLogger.Printf("%v added in stocks table.", s.Ticker)
}

// AddStocks from a trading platform
func AddStocks(s []Stock) {
	q := "INSERT INTO stocks(avanzaId, ticker, name) VALUES "
	vals := []interface{}{}
	for i := range s {
		s[i].Ticker = getTicker(s[i])
		q += "(?, ?, ?),"
		vals = append(vals, s[i].ID, s[i].Ticker, s[i].Name)
	}
	q = q[0:len(q)-1]
	if len(vals) > 0 {
		stmt := db.Prepare(q)
		stmt.Exec(vals...)
	}
}

// GetFirstNorth stocks from Avanza
func GetFirstNorth() {
	c := colly.NewCollector()
	stocks := []Stock{}
	c.OnHTML("tbody tr .orderbookName a[href]", func(e *colly.HTMLElement) {
		var stock Stock
		u := e.Attr("href")
		u = strings.TrimPrefix(u, "/aktier/om-aktien.html/")
		split := strings.Split(u, "/")
		stock.ID = split[0]
		stock.Name = split[1]
		stocks = append(stocks, stock)
	})
	url := "https://www.avanza.se/frontend/template.html/marketing/advanced-filter/advanced-filter-template" +
	"?1600622837794&widgets.marketCapitalInSek.filter.lower=&widgets.marketCapitalInSek.filter.upper=&widgets.marketCapitalInSek" +
	".active=true&widgets.stockLists.filter.list%5B0%5D=SE.FNSE&widgets.stockLists.active=true&widgets.numberOfOwners.filter.lower=" +
	"&widgets.numberOfOwners.filter.upper=&widgets.numberOfOwners.active=true&parameters.startIndex=0&parameters.maxResults=400&parameters" +
	".selectedFields%5B0%5D=LATEST&parameters.selectedFields%5B1%5D=DEVELOPMENT_TODAY&parameters.selectedFields%5B2%5D=DEVELOPMENT_ONE_YEAR" +
	"&parameters.selectedFields%5B3%5D=MARKET_CAPITAL_IN_SEK&parameters.selectedFields%5B4%5D=PRICE_PER_EARNINGS&parameters" +
	".selectedFields%5B5%5D=DIRECT_YIELD&parameters.selectedFields%5B6%5D=NBR_OF_OWNERS&parameters.selectedFields%5B7%5D=LIST"
	c.Visit(url)
	existingStocks := len(GetStocks())
	if existingStocks != len(stocks) {
		AddStocks(stocks)
		newStocks := len(stocks) - existingStocks
		InfoLogger.Printf("%v First North stocks added in stocks table.", newStocks)
	} else {
		InfoLogger.Println("No new stocks added.")
	}
}

// GetStocks from db
func GetStocks() []Stock {
	r := db.Query("SELECT avanzaId, ticker, name FROM stocks;")
	defer r.Close()
	stocks := []Stock{}
	for r.Next() {
		var id string
		var ticker string
		var name string
		var stock Stock
		r.Scan(&id, &ticker, &name)
		stock.ID = id
		stock.Ticker = ticker
		stock.Name = name
		stocks = append(stocks, stock) 
	}
	return stocks
}

func getTicker(s Stock) string {
	c := colly.NewCollector()
	c.OnHTML("head", func (e *colly.HTMLElement) {
		u := e.ChildText("title")
		re := regexp.MustCompile(`\((.*?)\)`)
		submatchall := re.FindAllString(u, -1)
		for _, element := range submatchall {
			element = strings.Trim(element, "(")
			element = strings.Trim(element, ")")
			s.Ticker = element
		}
	})
	url := "https://www.avanza.se/aktier/dagens-avslut.html/" + s.ID + "/" + s.Name
	c.Visit(url)
	return s.Ticker
}
