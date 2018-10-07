package organizations

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/pubcast/pubcast/data"
	"github.com/pubcast/pubcast/data/models"
	"github.com/pubcast/pubcast/handlers"
	"github.com/pubcast/pubcast/lib/activity"
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

	// Convert to an ActivityPub object
	url, err := url.Parse(handlers.GetFullHostname() + "/api/org/" + slug)
	actor := activity.NewOrganization(org.Name, url)

	// Turn the org into JSON
	bytes, err := json.Marshal(actor)

	// 500 because something went wrong marshaling the org
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success!
	w.Write(bytes)
}
