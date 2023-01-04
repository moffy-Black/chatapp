package db_connect_test

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://moffy:moffy0223@localhost:5432/chatuser?sslmode=disable"
)

func TestConnectDB(t *testing.T) {
	var db *sql.DB
	var err error

	db, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
