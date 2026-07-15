package db

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=password dbname=buyer-database sslmode=disable"

	DB, err = sql.Open("pgx", connStr)

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
