package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver       = "postgres"
	dbSourceString = "postgresql://root:secret@localhost:5432/simply_sacco?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSourceString)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
