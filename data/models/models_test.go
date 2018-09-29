package models

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/metapods/metapods/data"
	"github.com/stretchr/testify/assert"
)

// Runs before everything
func init() {
	data.RegisterTestDB()
}

func TestGetGroup(t *testing.T) {
	db, err := data.NewTestDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

func TestPutGroup(t *testing.T) {
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	slug, err := PutGroup(db, "hats and ;DROP TABLES", "<html>oh boy</html>")
	assert.Nil(t, err)

	group, err := GetGroup(db, slug)
	assert.Nil(t, err)

	assert.Equal(t, slug, group.Slug)
	assert.Equal(t, "hats and ;DROP TABLES", group.Name)
	assert.Equal(t, "<html>oh boy</html>", group.Note)
}
