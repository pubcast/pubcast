package shows

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/pubcast/pubcast/activity"
	"github.com/pubcast/pubcast/data"
	"github.com/pubcast/pubcast/data/models"
	"github.com/pubcast/pubcast/handlers"
)

// Get returns an Show
//
// Expects a `{slug}` url variable
// in the route: `/api/show/{slug}`
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// These arrise from the server setting up the routes incorrectly
	if vars == nil {
		http.Error(w, "Bad request", http.StatusInternalServerError)
		return
	}
	slug := vars["slug"]
	if slug == "" {
		http.Error(w, "Bad request, no slug in url", http.StatusInternalServerError)
		return
	}

	// Attempt to grab the Show
	show, err := models.GetShow(data.GetPool(), slug)

	// 500 because something went wrong with the database
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 404 because something we couldn't find the Show
	if show == nil {
		http.Error(w, slug+" does not exist on this server", http.StatusNotFound)
		return
	}

	// Convert to an ActivityPub object
	url, err := url.Parse(handlers.GetFullHostname() + "/api/show/" + slug)
	actor := activity.NewOrganization(show.Name, url)

	// Turn the Show into JSON
	bytes, err := json.Marshal(actor)

	// 500 because something went wrong marshaling the Show
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success!
	w.Write(bytes)
}
