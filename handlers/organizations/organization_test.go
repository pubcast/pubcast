package organizations

import (
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
	db, err := data.ConnectToTestDB()
	if err != nil {
		assert.NoError(t, err)
		return
	}
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

func TestGetOrganization(t *testing.T) {
	db, err := data.ConnectToTestDB()
	if err != nil {
		assert.NoError(t, err)
		return
	}
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
