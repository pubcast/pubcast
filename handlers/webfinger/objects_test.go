package webfinger

import (
	"encoding/json"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// Stolen from here:
// https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
func removeWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func TestLinkMarshalsCorrectly(t *testing.T) {
	bytes, err := json.Marshal(Link{Rel: "rel", Type: "type", HREF: "href"})

	assert.Nil(t, err)
	assert.Equal(t, string(bytes), `{"rel":"rel","type":"type","href":"href"}`)
}

func TestActorMarshalsCorrectly(t *testing.T) {
	expected := `{
		"subject": "acct:alice@my-example.com",
	
		"links": [
			{
				"rel": "self",
				"type": "application/activity+json",
				"href": "https://my-example.com/actor"
			}
		]
	}`

	bytes, err := json.Marshal(Actor{
		Subject: "acct:alice@my-example.com",
		Links: []Link{
			{Rel: "self", Type: "application/activity+json", HREF: "https://my-example.com/actor"},
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, removeWhitespace(string(bytes)), removeWhitespace(expected))
}
