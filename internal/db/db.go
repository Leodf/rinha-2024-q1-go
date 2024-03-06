package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var PG *sql.DB

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "admin"
	dbPassword = "123"
	dbName     = "rinha"
)

func Init() (*sql.DB, error) {
	host := dbHost
	if v, ok := os.LookupEnv("DB_HOST"); ok {
		host = v
	}
	uri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(time.Minute * 3)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	PG = db
	return db, nil
}
