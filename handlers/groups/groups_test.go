package groups

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pubcast/pubcast/activity"
	"github.com/pubcast/pubcast/data"
	"github.com/pubcast/pubcast/data/models"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func init() {
	data.SetupTestDB()
}

func TestGetGroupGives404s(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/group/{slug}", Get)

	// Try and expect a 404
	r := httptest.NewRequest("GET", "https://localhost:8080/org/i-dont-exist", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 404, w.Code)
}

func TestGetGroupGives500s(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Check if a route with no "{slug}" in the URL will return a 500
	router := mux.NewRouter()
	router.HandleFunc("/group/", Get) // Note that we don't have a {slug} here
	r := httptest.NewRequest("GET", "https://localhost:8080/group/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 500, w.Code)

	// Check if a route setup with "{something-other-than-slug}" will return a 500
	router = mux.NewRouter()
	router.HandleFunc("/group/{something-other}", Get) // Note that we don't have a {slug} here
	r = httptest.NewRequest("GET", "https://localhost:8080/group/boop", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 500, w.Code)
}

func TestGetGroup(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup a dummy org
	note := "foo"
	slug, err := models.PutGroup(db, "planet", note)
	assert.Equal(t, "planet", slug) // sanity
	assert.NoError(t, err)

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/api/group/{slug}", Get)

	// GET the /org
	r := httptest.NewRequest("GET", "https://localhost:8080/api/group/"+slug, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	// Expect a reasonable response
	assert.Equal(t, 200, w.Code)

	var org activity.Group
	err = json.Unmarshal(w.Body.Bytes(), &org)
	assert.NoError(t, err, w.Body.String()+" \n--failed to unmarshal")

	assert.Equal(t, slug, org.Name)
}

func TestCreateGroup(t *testing.T) {
	db := data.ConnectToTestDB(t)
	defer db.Close()

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/api/group", Create)

	// Setup request body
	body, err := json.Marshal(createGroupRequest{
		Name: "joe rogan podcast",
		Note: "a cool podcast group",
	})
	assert.NoError(t, err)

	// POST /api/org
	r := httptest.NewRequest("POST", "https://localhost:8080/api/group", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code)

	// Expect to see a slug in the response
	var response createGroupResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "joe-rogan-podcast", response.Slug)

	// Expect to see something in the database
	org, err := models.GetGroup(db, "joe-rogan-podcast")
	assert.NoError(t, err)
	assert.Equal(t, "joe-rogan-podcast", org.Slug)
	assert.Equal(t, "joe rogan podcast", org.Name)
	assert.Equal(t, "a cool podcast group", org.Note)
}
