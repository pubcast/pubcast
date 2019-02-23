package organizations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pubcast/pubcast/data"
	"github.com/pubcast/pubcast/data/models"
)

type createOrganizationRequest struct {
	Name string `json:"name"`
	Note string `json:"note"`
}
type createOrganizationResponse struct {
	Slug string `json:"slug"`
}

// Create adds an organization to the database
func Create(w http.ResponseWriter, r *http.Request) {
	// Parse the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	var body createOrganizationRequest
	json.Unmarshal(b, &body)

	// Insert into db
	db := data.GetPool()
	slug, err := models.PutOrganization(db, body.Name, body.Note)
	if err != nil {
		http.Error(w, "Something went wrong inserting into the db", http.StatusInternalServerError)
		return
	}

	// Return slug in json
	response := createOrganizationResponse{Slug: slug}
	jsonstr, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Something went wrong responding from request", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonstr)
}
