package shows

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/pubcast/pubcast/activity"
	"github.com/pubcast/pubcast/data"
	"github.com/pubcast/pubcast/data/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	data.SetupTestDB()
}

func TestGetShowGives404s(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/org/{slug}", Get)

	// Try and expect a 404
	r := httptest.NewRequest("GET", "https://localhost:8080/show/i-dont-exist", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 404, w.Code)
}

func TestGetShowGives500s(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Check if a route with no "{slug}" in the URL will return a 500
	router := mux.NewRouter()
	router.HandleFunc("/show/", Get) // Note that we don't have a {slug} here
	r := httptest.NewRequest("GET", "https://localhost:8080/show/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 500, w.Code)

	// Check if a route setup with "{something-other-than-slug}" will return a 500
	router = mux.NewRouter()
	router.HandleFunc("/show/{something-other}", Get) // Note that we don't have a {slug} here
	r = httptest.NewRequest("GET", "https://localhost:8080/show/boop", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 500, w.Code)
}

func TestGetShow(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup a dummy show
	note := "foo"
	slug, err := models.PutShow(db, "planet", note)
	assert.Equal(t, "planet", slug) // sanity
	assert.NoError(t, err)

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/api/show/{slug}", Get)

	// GET the /show
	r := httptest.NewRequest("GET", "https://localhost:8080/api/show/"+slug, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	// Expect a reasonable response
	assert.Equal(t, 200, w.Code)

	var show activity.Organization
	err = json.Unmarshal(w.Body.Bytes(), &show)
	assert.NoError(t, err, w.Body.String()+" \n--failed to unmarshal")

	assert.Equal(t, slug, show.Name)
}

func TestCreateShow(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/api/show", Create)

	// Setup request body
	body, err := json.Marshal(createShowRequest{
		Name: "jims cool podcasts",
		Note: "a cool podcast show",
	})
	assert.NoError(t, err)

	// POST /api/show
	r := httptest.NewRequest("POST", "https://localhost:8080/api/show", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code)

	// Expect to see a slug in the response
	var response createShowResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "jims-cool-podcasts", response.Slug)

	// Expect to see something in the database
	show, err := models.GetShow(db, "jims-cool-podcasts")
	assert.NoError(t, err)
	assert.Equal(t, "jims-cool-podcasts", show.Slug)
	assert.Equal(t, "jims cool podcasts", show.Name)
	assert.Equal(t, "a cool podcast show", show.Note)
}
