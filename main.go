package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/trungvdn/simplebank/api"
	db "github.com/trungvdn/simplebank/db/sqlc"
	"github.com/trungvdn/simplebank/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}

	db := db.NewStore(conn)
	server := api.NewServer(db)
	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
