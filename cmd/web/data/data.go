package data

import (
	"github.com/jmoiron/sqlx"
	"os"
)

var connectionKey string = "MYSQL_CONNECTION"
var connectionString string

// db .
var db *sqlx.DB

func init() {
	connectionString = os.Getenv(connectionKey)
	if connectionString == "" {
		panic("ENV: " + connectionKey + " does not exist.")
	}

	sqlxdb, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	db = sqlxdb
}
