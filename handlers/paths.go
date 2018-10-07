package handlers

import (
	"github.com/metapods/metapods/config"
	"github.com/spf13/viper"
)

// GetFullHostname returns something like `https://foobang.com:8080`
func GetFullHostname() string {
	return "https://" + viper.GetString(config.ServerHostname) +
		":" + viper.GetString(config.ServerPort)
}
