package config

import (
	"database/sql"
	"log"
	// "os"

	db "github.com/goweb3/services/database/sql"
)

// DB func
func DB() *sql.DB {
	// dbDlct := os.Getenv("DATABASE_DLCT")
	// dbUser := os.Getenv("DATABASE_USER")
	// dbPass := os.Getenv("DATABASE_PASS")
	// dbHost := os.Getenv("DATABASE_HOST")
	// dbPort := os.Getenv("DATABASE_PORT")
	// dbName := os.Getenv("DATABASE_NAME")
	dbDlct := "postgres"
	dbUser := "default"
	dbPass := "sceret"
	dbHost := "127.0.0.1"
	dbPort := "5432"
	dbName := "default"
	db, err := db.Connect(dbDlct, dbUser, dbPass, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
