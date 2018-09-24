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

	value := application.Owns(application{}, context.TODO(), u)
	assert.True(t, value, "Owns should returns true if the url matches the config")
}

func TestOwnsFailsWithWrongHostname(t *testing.T) {

	// We use set default to mock the required config info
	viper.SetDefault(config.ServerHostname, "awrongsite.com")
	viper.SetDefault(config.ServerPort, 8080)

	u, _ := url.Parse("https://goodboys.com:8080")

	value := application.Owns(application{}, context.TODO(), u)
	assert.False(t, value, "Owns returns false if the url doesnt match the config")
}
