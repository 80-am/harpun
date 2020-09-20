package cmd

import (
	"github.com/80-am/harpun/db"
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