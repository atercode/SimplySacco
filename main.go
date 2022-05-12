package main

import (
	"database/sql"
	"log"

	"github.com/atercode/SimplySacco/api"
	db "github.com/atercode/SimplySacco/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver       = "postgres"
	dbSourceString = "postgresql://root:secret@localhost:5432/simply_sacco?sslmode=disable"
	serverAddress  = "localhost:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSourceString)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
