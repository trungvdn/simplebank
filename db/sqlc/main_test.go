package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/trungvdn/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
