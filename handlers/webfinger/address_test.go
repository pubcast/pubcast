package webfinger

import (
	"testing"

	_ "github.com/lib/pq"

	"github.com/metapods/metapods/data"
	"github.com/stretchr/testify/assert"
)

// Runs before everything
func init() {
	data.SetupTestDB()
	data.ConnectToTestDB()
}

func TestAtAddressCanFail(t *testing.T) {
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
