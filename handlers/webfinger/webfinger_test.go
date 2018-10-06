package webfinger

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/metapods/metapods/config"
	"github.com/metapods/metapods/data"
	"github.com/metapods/metapods/data/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	data.SetupTestDB()
}

func TestWebfingerBadRequests(t *testing.T) {
	db, err := data.ConnectToTestDB()
	if err != nil {
		assert.NoError(t, err)
		return
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
		assert.NoError(t, err)
		return
	}
	defer db.Close()

	// Setup Config
	viper.SetDefault(config.ServerHostname, "localhost")
	viper.SetDefault(config.ServerPort, "8080")

	// Setup a dummy organization
	slug, err := models.PutOrganization(db, "slurp", "bloop")
	assert.Equal(t, "slurp", slug) // sanity test
	assert.NoError(t, err)

	// Query they webfinger endpoint
	query := "?resource=acct:" + slug + "@fooman.org"
	r := httptest.NewRequest("GET",
		"https://localhost:8080/.well-known/webfinger"+query, nil)
	w := httptest.NewRecorder()
	Get(w, r)

	// Expect a _valid_ response (good status code, correct format, etc)
	assert.Equal(t, 200, w.Code)
	assert.True(t, len(w.Body.String()) > 0)

	// Expect a _correct_ response (actually returns the same org that we put in)
	var actor Actor
	err = json.Unmarshal(w.Body.Bytes(), &actor)
	assert.NoError(t, err)
	assert.Equal(t, slug+"@fooman.org", actor.Subject)
	assert.Equal(t, "https://localhost:8080/activity/organizations/slurp", actor.Links[0].HREF)
	assert.Equal(t, "self", actor.Links[0].Rel)
	assert.Equal(t, "application/activity+json", actor.Links[0].Type)
}
