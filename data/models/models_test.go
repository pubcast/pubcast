package models

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
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
	txdb.Register("txdb", "postgres", psqlInfo)
}

func TestGetGroup(t *testing.T) {
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}

	// Populate the db with some dummy data
	query := `
		INSERT INTO groups (slug, name, note)
		VALUES ('dog', 'Corgies', 'I like pups')
	`
	_, err = db.Exec(query)
	assert.Nil(t, err) // Inserts should succeed

	group, err := GetGroup(db, "dog")
	assert.Nil(t, err)
	assert.NotNil(t, group)

	assert.Equal(t, "dog", group.Slug, "Group slug should match")
	assert.Equal(t, "Corgies", group.Name, "Group name should match")
	assert.Equal(t, "I like pups", group.Note, "Group note should match")
}
