package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/ah98lg/al_bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfiguration("../..")
	if err != nil {
		log.Fatal("Could not load config file: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
