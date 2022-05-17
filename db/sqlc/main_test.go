package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/atercode/SimplySacco/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load configs: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSourceSting)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
