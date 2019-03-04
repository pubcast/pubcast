package webfinger

import (
	"testing"

	_ "github.com/lib/pq"

	"github.com/pubcast/pubcast/data"
	"github.com/pubcast/pubcast/data/models"
	"github.com/stretchr/testify/assert"
)

// Runs before everything
func init() {
	data.SetupTestDB()
}

func TestAtAddressCanFail(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	var tests = []struct {
		id     string
		passes bool
	}{
		// Good data
		{"joe@foodog.com", true},
		{"foo@marco.polo.edu", true},
		{"sf.county@marco.polo.org", true},
		{"w@f.c", true},

		// Bad data
		{" ", false},
		{"@", false},
		{"m@", false},
		{"not-an-address", false},
	}

	for _, row := range tests {
		_, err := atAddress(row.id)

		var message string
		if row.passes {
			message = " should have passed"
		} else {
			message = " should have failed"
		}

		assert.Equal(t, row.passes, err == nil, row.id+message)
	}
}

func TestAtAddressReturnsNilIfNoAddress(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	actor, err := atAddress("m@iscool.org")
	assert.Nil(t, err)
	assert.Nil(t, actor)
}

func TestAtAddressReturnsShowReferenceIfExists(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	slug, err := models.PutShow(db, "woo", "a note")
	assert.Equal(t, "woo", slug) // Sanity check

	actor, err := atAddress("woo@moo.org")
	assert.Nil(t, err)
	assert.NotNil(t, actor)
	assert.Equal(t, "woo@moo.org", actor.Subject)
}
