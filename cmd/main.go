package cmd

import (
	"flag"

	"github.com/80-am/harpun/db"
)

var c Config
var initDb bool

func init() {
    flag.BoolVar(&initDb, "initDb", false, "Initializes your stocks table with First North Stockholm data")
    flag.Parse()
}

func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}

// Main for harpun app
func Main() {
	c.GetConfig()
	database, err := db.Init(c.DbUser, c.DbPassword, c.DbSchema)
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()

	if isFlagPassed("initDb") {
		GetFirstNorth()
	}
	stocks := GetStocks()
	for _ , stock := range stocks {
		GetDailyTrades(stock)
	}
}
