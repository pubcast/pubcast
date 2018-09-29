package data

import (
	"database/sql"
	"fmt"

	// TODO: move this to a main package
	txdb "github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
)

var (
	// Pool is the pool opened for the database
	Pool *sql.DB
)

// NewDB opens a standard DB
func NewDB() (*sql.DB, error) {

	const (
		host   = "localhost"
		port   = 5432
		user   = "postgres"
		dbname = "metapods"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	return sql.Open("postgres", psqlInfo)
}

// RegisterTestDB is used to setup a transactional database.
// Use it inside of an `init` function in a test file.
func RegisterTestDB() {
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
	txdb.Register("txdb", "postgres", psqlInfo)
}

// NewTestDB creates a new of the test database
func NewTestDB() (*sql.DB, error) {
	return sql.Open("txdb", "identifier")
}

// InitNewTestDB creates a new test db pool and sets it to data.Pool
// Call this if you're using data.Pool somewhere and want your test
// to use our test db.
//
// If you're trying to use this and
func InitNewTestDB() (*sql.DB, error) {
	db, err := NewTestDB()
	if err != nil {
		return db, err
	}

	Pool = db
	return db, nil
}
