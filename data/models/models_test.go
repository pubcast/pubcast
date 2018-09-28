package models

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
)

// Runs before everything
func init() {
	const (
		host   = "localhost"
		port   = 5432
		user   = "postgres"
		dbname = "metapods_test"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	// we register an sql driver named "txdb"
	txdb.Register("txdb", "pq", psqlInfo)
}

func TestGetGroup(t *testing.T) {
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
