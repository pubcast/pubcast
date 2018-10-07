package handlers

import (
	"strings"
	"testing"

	"github.com/metapods/metapods/config"
	"github.com/stretchr/testify/assert"

	"github.com/spf13/viper"
)

func TestGetFullHostnameUsesHttps(t *testing.T) {
	viper.SetDefault(config.ServerHostname, "foo")
	viper.SetDefault(config.ServerPort, 8080)

	path := GetFullHostname()
	assert.True(t, strings.HasPrefix(path, "https://"), "Must support https://")
}
