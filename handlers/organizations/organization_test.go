package organizations

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

func TestGetOrganizationGives404s(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/org/{slug}", Get)

	// Try and expect a 404
	r := httptest.NewRequest("GET", "https://localhost:8080/org/i-dont-exist", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 404, w.Code)
}

func TestGetOrganizationGives500s(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Check if a route with no "{slug}" in the URL will return a 500
	router := mux.NewRouter()
	router.HandleFunc("/org/", Get) // Note that we don't have a {slug} here
	r := httptest.NewRequest("GET", "https://localhost:8080/org/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 500, w.Code)

	// Check if a route setup with "{something-other-than-slug}" will return a 500
	router = mux.NewRouter()
	router.HandleFunc("/org/{something-other}", Get) // Note that we don't have a {slug} here
	r = httptest.NewRequest("GET", "https://localhost:8080/org/boop", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 500, w.Code)
}

func TestGetOrganization(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup a dummy org
	note := "foo"
	slug, err := models.PutOrganization(db, "planet", note)
	assert.Equal(t, "planet", slug) // sanity
	assert.NoError(t, err)

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/api/org/{slug}", Get)

	// GET the /org
	r := httptest.NewRequest("GET", "https://localhost:8080/api/org/"+slug, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	// Expect a reasonable response
	assert.Equal(t, 200, w.Code)

	var org activity.Organization
	err = json.Unmarshal(w.Body.Bytes(), &org)
	assert.NoError(t, err, w.Body.String()+" \n--failed to unmarshal")

	assert.Equal(t, slug, org.Name)
}

func TestCreateOrganization(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/api/org", Create)

	// Setup request body
	body, err := json.Marshal(createOrganizationRequest{
		Name: "jims cool podcasts",
		Note: "a cool podcast group",
	})
	assert.NoError(t, err)

	// POST /api/org
	r := httptest.NewRequest("POST", "https://localhost:8080/api/org", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code)

	// Expect to see a slug in the response
	var response createOrganizationResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "jims-cool-podcasts", response.Slug)

	// Expect to see something in the database
	org, err := models.GetOrganization(db, "jims-cool-podcasts")
	assert.NoError(t, err)
	assert.Equal(t, "jims-cool-podcasts", org.Slug)
	assert.Equal(t, "jims cool podcasts", org.Name)
	assert.Equal(t, "a cool podcast group", org.Note)
}
