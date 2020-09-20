package cmd

import (
	"fmt"
	"strconv"
    "strings"


	"github.com/80-am/harpun/db"
	"github.com/gocolly/colly/v2"
)

// Trade fetched for a ticker
type Trade struct {
	Ticker string
	Buyer  string
	Seller string
	Amount float64
	Price  float64
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
		t.Amount, _ = strconv.ParseFloat(strings.Replace(e.ChildText(":nth-child(3)"), ",", ".", -1), 8)
		t.Price, _ = strconv.ParseFloat(strings.Replace(e.ChildText(":nth-child(4)"), ",", ".", -1), 8)
		t.Time = e.ChildText(".last")
		if t.Time != "" {
			trades = append(trades, t)
		}
	})
	url := "https://www.avanza.se/aktier/dagens-avslut.html/" + s.ID + "/" + s.Ticker
	c.Visit(url)
	updateTrades(trades)
}

func updateTrades(t []Trade) {
	q := "INSERT INTO trades(ticker, buyer, seller, amount, price, time) VALUES "
	vals := []interface{}{}

	for i := range t {
		fmt.Println(t[i].Amount, t[i].Price)
	    q += "(?, ?, ?, ?, ?, ?),"
		vals = append(vals, t[i].Ticker, t[i].Buyer, t[i].Seller, t[i].Amount, t[i].Price, t[i].Time)
	}
	q = q[0:len(q)-1]
	stmt := db.Prepare(q)
	stmt.Exec(vals...)
}