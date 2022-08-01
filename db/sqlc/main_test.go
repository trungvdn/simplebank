package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	driverName = "postgres"
	sourceName = "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driverName, sourceName)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
