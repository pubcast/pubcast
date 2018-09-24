package activitypub

import (
	"context"
	"net/url"
	"testing"

	"github.com/Flaque/metapod/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestOwnsCanConformToTheConfig(t *testing.T) {

	// We use set default to mock the required config info
	viper.SetDefault(config.ServerHostname, "bestjeanist.com")
	viper.SetDefault(config.ServerPort, 8080)

	u, _ := url.Parse("https://bestjeanist.com:8080")

	doesOwn := application.Owns(application{}, context.TODO(), u)
	assert.True(t, doesOwn, "Owns should returns true if the url matches the config")
}

func TestOwnsFailsWithWrongHostname(t *testing.T) {

	// We use set default to mock the required config info
	viper.SetDefault(config.ServerHostname, "awrongsite.com")
	viper.SetDefault(config.ServerPort, 8080)

	u, _ := url.Parse("https://goodboys.com:8080")

	doesOwn := application.Owns(application{}, context.TODO(), u)
	assert.False(t, doesOwn, "Owns returns false if the url doesnt match the config")
}

func TestHas(t *testing.T) {
	var tests = []struct {
		path     string
		expected bool
	}{
		// Good data
		{"/activity/organization/joey/inbox", true},
		{"/activity/organization/joey/outbox", true},
		{"/activity/organization/sam/inbox", true},

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
		doesMatch, err := application.Has(application{}, context.TODO(), u)
		assert.Nil(t, err, "Has should not fail")
		assert.Equal(t, row.expected, doesMatch, row.path+" failed")
	}
}
