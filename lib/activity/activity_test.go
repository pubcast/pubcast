package activity

import (
	"log"
	"net/url"
	"testing"

	_ "github.com/DATA-DOG/go-txdb"
	"github.com/metapods/metapods/config"
	"github.com/metapods/metapods/data"
	"github.com/metapods/metapods/data/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	data.SetupTestDB()
}

func TestOwnsCanConformToTheConfig(t *testing.T) {

	// We use set default to mock the required config info
	viper.SetDefault(config.ServerHostname, "bestjeanist.com")
	viper.SetDefault(config.ServerPort, 8080)

	u, _ := url.Parse("https://bestjeanist.com:8080")

	doesOwn := Owns(u)
	assert.True(t, doesOwn, "Owns should returns true if the url matches the config")
}

func TestOwnsFailsWithWrongHostname(t *testing.T) {

	// We use set default to mock the required config info
	viper.SetDefault(config.ServerHostname, "awrongsite.com")
	viper.SetDefault(config.ServerPort, 8080)

	u, _ := url.Parse("https://goodboys.com:8080")

	doesOwn := Owns(u)
	assert.False(t, doesOwn, "Owns returns false if the url doesnt match the config")
}

func TestMatchesURLSpec(t *testing.T) {
	var tests = []struct {
		path     string
		expected bool
	}{
		// Good data
		{"/activity/organization/joey/inbox", true},
		{"/activity/organization/joey/outbox", true},
		{"/activity/organization/sam/inbox", true},
		{"/activity/group/joey/inbox", true},
		{"/activity/group/joey/outbox", true},
		{"/activity/group/sam/inbox", true},

		// Bad data
		{"/activity/organization/joey/something", false},
		{"/activity/bad/joey/inbox", false},
		{"/something/organization/joey/inbox", false},
		{"/activity", false},
		{"/activity/organization/", false},
		{"/activity/organization/joey/", false},
	}

	var base = "https://podocasto.com:8080"

	for _, row := range tests {
		u, _ := url.Parse(base + row.path)
		doesMatch := matchesURLSpec(u)
		assert.Equal(t, row.expected, doesMatch, row.path+" failed")
	}
}

func TestHasFailsIfNothingExists(t *testing.T) {

	db, err := data.ConnectToTestDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var base = "https://podocasto.com:8080"
	u, _ := url.Parse(base + "/activity/group/sam/inbox")

	val, err := Has(u)
	assert.Nil(t, err)
	assert.False(t, val)
}

func TestHasPassesIfSomethingExists(t *testing.T) {
	db, err := data.ConnectToTestDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Dummy data
	slug, err := models.PutGroup(db, "joeboe", "woo")
	assert.Nil(t, err)

	var base = "https://podocasto.com:8080"
	u, _ := url.Parse(base + "/activity/group/" + slug + "/inbox")

	val, err := Has(u)
	assert.Nil(t, err)
	assert.True(t, val)
}
