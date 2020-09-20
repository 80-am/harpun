package db

import (
    "fmt"
    "database/sql"

    // Used for initialization
    _ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

// Init the database connection
func Init(user string, pass string, schema string) (*sql.DB, error) {
    dbLink, err := sql.Open("mysql", user+":"+pass+"@"+schema)
    if err != nil {
        return nil, err 
    }   
    err = dbLink.Ping()
    if err != nil {
        return nil, err 
    }   

    database = dbLink
    return database, nil 
}

// Query to mysql
func Query(query string) *sql.Rows {
    rows, err := database.Query(query)
    if err != nil {
        fmt.Println(err)
    }   
    return rows
}

// QueryRow to mysql
func QueryRow (query string, args interface{}) *sql.Row {
    row := database.QueryRow(query, args)
    return row 
}

// Prepare query to be executed
func Prepare(query string) *sql.Stmt {
    stmt, err := database.Prepare(query)
    if err != nil {
        fmt.Println(err)
    }   
    return stmt
}
