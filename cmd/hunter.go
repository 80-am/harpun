package cmd

import (
	"math"
	"strconv"
	"time"

	"github.com/80-am/harpun/db"
)

// DetectWhale aka unusual large trade
func DetectWhale(s Stock, t Trade) {
	avgAmount := getAverageAmount(t.Ticker)
	if avgAmount == 0 {
		return
	}
	unusualAmount := avgAmount * Multiplier
	if t.Amount > unusualAmount {
		price, _ := strconv.ParseFloat(t.Price, 8)
		totalPrice := math.Round((float64(t.Amount) * price) * 100) / 100
		alert(s, t, totalPrice)
	}
}

// UpdateAverageAmounts for stocks
func UpdateAverageAmounts(s Stock) {
	q := "SELECT amount FROM trades WHERE ticker = '" + s.Ticker + "' ORDER BY amount DESC LIMIT 100;"
	r := db.Query(q)
	defer r.Close()
	avgAmount := getAverageAmount(s.Ticker)
	amounts := []int64{}
	for r.Next() {
		var amount int64
		r.Scan(&amount)
		amounts = append(amounts, amount)
	}
	var total int64 = 0
	for _, amount := range amounts {
		total += amount
	}
	if total > 0 {
		avgTradeAmount := total / int64(len(amounts))
		if avgAmount != avgTradeAmount {
			stmt := db.Prepare("UPDATE stocks SET avgTradeAmount = ? WHERE ticker = ?;")
			stmt.Exec(avgTradeAmount, s.Ticker)
			InfoLogger.Printf("New trade average %v for %v", avgTradeAmount, s.Ticker)
		}
	}
}

func alert(s Stock, t Trade, tp float64) {
	r := db.QueryRow("SELECT time FROM alerts WHERE ticker = (?) ORDER BY time DESC LIMIT 1;", t.Ticker)
	var lastAlert string
	r.Scan(&lastAlert)
	date := time.Now().Format("2006-01-02")
	dateTime := date + " " + t.Time
	if lastAlert != dateTime {
		stmt := db.Prepare("INSERT INTO alerts(ticker, amount, price, totalPrice, time) VALUES (?, ?, ?, ?, ?)")
		stmt.Exec(t.Ticker, t.Amount, t.Price, tp, dateTime)
		AlertLogger.Printf("%s %v shares bought @ %s for a total of %.2f SEK (Trade time: %s)\n", t.Ticker, t.Amount, t.Price, tp, dateTime)
		if (Hook) {
			AlertHook(s, t, tp)
		}
	}
}

func getAverageAmount(t string) int64 {
	r := db.QueryRow("SELECT avgTradeAmount FROM stocks WHERE ticker = (?);", t)
	var avgAmount int64
	r.Scan(&avgAmount)
	return avgAmount
}