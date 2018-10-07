package organizations

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pubcast/pubcast/data"
	"github.com/pubcast/pubcast/data/models"
)

// Get returns an Organization
//
// Expects a `{slug}` url variable
// in the route: `/api/org/{slug}`
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Handle 400s
	if vars == nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	slug := vars["slug"]
	if slug == "" {
		http.Error(w, "Bad request, no slug in url", http.StatusBadRequest)
		return
	}

	// Attempt to grab the org
	org, err := models.GetOrganization(data.GetPool(), slug)

	// 500 because something went wrong with the database
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 404 because something we couldn't find the organization
	if org == nil {
		http.Error(w, slug+" does not exist on this server", http.StatusNotFound)
		return
	}

	// Turn the org into JSON
	bytes, err := json.Marshal(org)

	// 500 because something went wrong marshaling the org
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success!
	w.Write(bytes)
}
