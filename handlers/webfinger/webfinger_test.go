package webfinger

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebfinger(t *testing.T) {

	r := httptest.NewRequest("GET", "https://localhost:8080/.well-known/webfinger", nil)
	w := httptest.NewRecorder()

	Get(w, r)

	assert.Equal(t, w.Code, 200)
	assert.NotEqual(t, w.Body.String(), "")
}
