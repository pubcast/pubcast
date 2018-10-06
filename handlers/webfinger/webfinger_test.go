package webfinger

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/metapods/metapods/data"
	"github.com/metapods/metapods/data/models"
	"github.com/stretchr/testify/assert"
)

func TestWebfingerBadRequests(t *testing.T) {
	db, err := data.ConnectToTestDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var tests = []struct {
		query string
		code  int
	}{
		// 400s
		{"?not-resource", http.StatusBadRequest},
		{"?resource", http.StatusBadRequest},
		{"?resource=::foomandoop@bloop", http.StatusBadRequest},
		{"?resource=acct:not-an-address", http.StatusBadRequest},

		// 404s
		{"?resource=acct:joe@moo.org", http.StatusNotFound},
	}

	for _, test := range tests {
		r := httptest.NewRequest("GET",
			"https://localhost:8080/.well-known/webfinger"+test.query, nil)
		w := httptest.NewRecorder()

		Get(w, r)

		assert.Equal(t, test.code, w.Code, test.query+" should have a status of "+strconv.Itoa(test.code))
	}
}

func TestWebfingerSuccessfulRequest(t *testing.T) {
	db, err := data.ConnectToTestDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	slug, err := models.PutOrganization(db, "slurp", "bloop")
	assert.Equal(t, "slurp", slug) // sanity test
	assert.Nil(t, err)

	query := "?resource=acct:" + slug + "@fooman.org"

	r := httptest.NewRequest("GET",
		"https://localhost:8080/.well-known/webfinger"+query, nil)
	w := httptest.NewRecorder()

	Get(w, r)

	assert.Equal(t, 200, w.Code)

	var org models.Organization
	json.Unmarshal(w.Body.Bytes(), &org)

}
