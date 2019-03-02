package activitypub

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestInbox(t *testing.T) {
	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/inbox", Inbox)

	// Try and expect a 200
	r := httptest.NewRequest("POST", "https://localhost:8080/inbox", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code)
}
