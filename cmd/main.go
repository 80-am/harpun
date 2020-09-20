package cmd

import (
	"github.com/80-am/harpun/db"
)

var c Config
var s Stock

// Main for harpun app
func Main() {
	c.GetConfig()
	database, err := db.Init(c.DbUser, c.DbPassword, c.DbSchema)
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()
	
	stocks := GetStocks()
	for _ , stock := range stocks {
		GetDailyTrades(stock)
	}
}
