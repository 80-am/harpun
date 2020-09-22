package cmd

import (
	"strconv"
	"strings"
	"time"

	"github.com/80-am/harpun/db"
	"github.com/gocolly/colly/v2"
)

// Trade fetched for a ticker
type Trade struct {
	Ticker string
	Buyer  string
	Seller string
	Amount int64
	Price  string
	Time   string
}

// GetDailyTrades for a ticker, current limit is 100 trades per fetch
func GetDailyTrades(s Stock) {
	c := colly.NewCollector()
	trades := []Trade{}
	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		var t Trade
		t.Ticker = s.Ticker
		t.Buyer = e.ChildText(".tLeft:nth-child(1)")
		t.Seller = e.ChildText(".tLeft:nth-child(2)")
		a := e.ChildText(":nth-child(3)")
		t.Amount, _ = strconv.ParseInt(strings.Join(strings.Fields(strings.TrimSpace(a)), ""),10, 64)
		t.Price = strings.Replace(e.ChildText(":nth-child(4)"), ",", ".", -1)
		t.Time = e.ChildText(".last")
		if t.Time != "" {
			trades = append(trades, t)
		}
	})
	url := "https://www.avanza.se/aktier/dagens-avslut.html/" + s.ID + "/" + s.Name
	c.Visit(url)
	if len(trades) > 0 {
		updateTrades(trades)
	}
}

func getLastTrade(t Trade) string {
	r := db.QueryRow("SELECT time FROM trades WHERE ticker = (?) ORDER BY time DESC LIMIT 1;", t.Ticker)
	var lastTrade string
	r.Scan(&lastTrade)
	return lastTrade
}

func updateTrades(t []Trade) {
	q := "INSERT INTO trades(ticker, buyer, seller, amount, price, time) VALUES "
	vals := []interface{}{}
	lastTrade := getLastTrade(t[0])
	date := time.Now().Format("2006-01-02")
	newTrades := 0
	for i := len(t)-1; i >= 0; i-- {
		dateTime := date + " " + t[i].Time
		if dateTime > lastTrade {
			q += "(?, ?, ?, ?, ?, ?),"
			vals = append(vals, t[i].Ticker, t[i].Buyer, t[i].Seller, t[i].Amount, t[i].Price, dateTime)
			newTrades++
		}
	}
	q = q[0:len(q)-1]
	if len(vals) > 0 {
		stmt := db.Prepare(q)
		stmt.Exec(vals...)
	}
	if newTrades > 0 {
		InfoLogger.Printf("Collected %v trades for %v.", newTrades, t[0].Ticker)
	}
}