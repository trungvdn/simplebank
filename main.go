package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/trungvdn/simplebank/api"
	db "github.com/trungvdn/simplebank/db/sqlc"
)

const (
	driverName = "postgres"
	sourceName = "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable"
	address    = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(driverName, sourceName)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}

	db := db.NewStore(conn)
	server := api.NewServer(db)
	if err = server.Start(address); err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
