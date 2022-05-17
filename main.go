package main

import (
	"database/sql"
	"log"

	"github.com/atercode/SimplySacco/api"
	db "github.com/atercode/SimplySacco/db/sqlc"
	"github.com/atercode/SimplySacco/utils"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configs: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSourceSting)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
