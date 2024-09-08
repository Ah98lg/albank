package main

import (
	"database/sql"
	"log"

	"github.com/ah98lg/al_bank/api"
	db "github.com/ah98lg/al_bank/db/sqlc"
	"github.com/ah98lg/al_bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfiguration(".")
	if err != nil {
		log.Fatal("Could not load config file", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
