package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/80-am/harpun/db"
)

var c Config
var initDb bool
var config string
// Multiplier defining large trades
var Multiplier int64
// WarningLogger for harpun
var WarningLogger *log.Logger
// InfoLogger for harpun
var InfoLogger *log.Logger
// ErrorLogger for harpun
var ErrorLogger *log.Logger
// AlertLogger for harpun
var AlertLogger *log.Logger 

func init() {
	flag.BoolVar(&initDb, "initDb", false, "Initializes your stocks table with First North Stockholm data")
	flag.StringVar(&config, "config", "", "Path to your config.yml")
	flag.Parse()
	file, err := os.OpenFile("harpun.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
	}
	InfoLogger = log.New(file, "[INFO] ", log.Ldate|log.Ltime)
	WarningLogger = log.New(file, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	AlertLogger = log.New(file, "[ALERT] ", log.Ldate|log.Ltime)
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
	InfoLogger.Println("Harpun started.")
	c.GetConfig(config)
	if c.Multiplier != 0 {
		Multiplier = c.Multiplier
	} else {
		Multiplier = 5
	}
	database, err := db.Init(c.DbUser, c.DbPassword, c.DbSchema)
	if err != nil {
		ErrorLogger.Println(err.Error())
		panic(err.Error())
	}
	defer database.Close()

	if isFlagPassed("initDb") {
		GetFirstNorth()
	}
	stocks := GetStocks()
	for _ , stock := range stocks {
		GetDailyTrades(stock)
		UpdateAverageAmounts(stock)
	}
}
