package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("DB not reachable: %v", err)
	}
	return db
}
