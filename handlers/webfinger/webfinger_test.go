package webfinger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/pubcast/pubcast/activity"
	"github.com/pubcast/pubcast/config"
	"github.com/pubcast/pubcast/data"
	"github.com/pubcast/pubcast/data/models"
	"github.com/pubcast/pubcast/handlers/shows"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	data.SetupTestDB()
}

func TestWebfingerBadRequests(t *testing.T) {
	db := data.ConnectToTestDB(t)
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

// Tests the _entirety_ of a successful webfinger request.
func TestWebfingerSuccessfulRequest(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup Config
	viper.SetDefault(config.ServerHostname, "localhost")
	viper.SetDefault(config.ServerPort, "8080")

	// Setup a dummy org/show
	note := "bloop"
	slug, err := models.PutShow(db, "slurp", note)
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
	assert.Equal(t, "https://localhost:8080/api/show/slurp", actor.Links[0].HREF)
	assert.Equal(t, "self", actor.Links[0].Rel)
	assert.Equal(t, "application/activity+json", actor.Links[0].Type)

	// Now check that correct response to also be queryable
	// This effectively completes the handshake and assures
	// our interoperability with other services.
	r = httptest.NewRequest("GET", actor.Links[0].HREF, nil)
	w = httptest.NewRecorder()
	router := mux.NewRouter()
	fmt.Println(actor.Links[0].HREF)
	router.HandleFunc("/api/show/{slug}", shows.Get)
	router.ServeHTTP(w, r)

	// Expect a reasonable response from the show
	assert.Equal(t, 200, w.Code)

	// Finally, check that we can correctly get an org/show
	var org activity.Organization
	err = json.Unmarshal(w.Body.Bytes(), &org)
	assert.NoError(t, err, w.Body.String()+" \n--failed to unmarshal")
	assert.Equal(t, slug, org.Name)
	assert.Equal(t, "https://www.w3.org/ns/activitystreams", org.Context)
}
