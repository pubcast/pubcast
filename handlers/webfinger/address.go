/*
An address is a wrapper around the Organization actor type.

To interop with the rest of the world, Organization's are given identifiers that look like this: "planet-money@ex-host-server.org".

Combined with webfinger, this is the outside world's way of finding a particular resource on our server.
*/

package webfinger

import (
	"net/mail"
	"strings"

	"github.com/pubcast/pubcast/config"
	"github.com/pubcast/pubcast/data"
	"github.com/pubcast/pubcast/data/models"
	"github.com/spf13/viper"
)

type badAddressError struct {
	address string
}

func (e *badAddressError) Error() string {
	return "badly formatted address: " + e.address
}

func newBadAddressError(address string) *badAddressError {
	return &badAddressError{address: address}
}

// slugOf gets the "slug" of an address; which is just the part to the left of the @.
// So for `planet-money@foo.org`, the slug would be `planet-money`.
func slugOf(address string) string {
	fragments := strings.Split(address, "@")
	return fragments[0]
}

func atAddress(address string) (*Actor, error) {

	// Although these aren't technically email addresses,
	// they still match a subset of the RFC 5322 email format
	parser := mail.AddressParser{}
	_, err := parser.Parse(address)
	if err != nil {
		return nil, newBadAddressError(address)
	}

	// 99pi@blah.org => 99pi
	slug := slugOf(address)

	// If you're thinking:
	// > "Isn't this really inefficient to get the object
	//   and then just return a reference so we can get it again?"
	//
	// You would be right.
	org, err := models.GetOrganization(data.GetPool(), slug)
	if err != nil {
		return nil, err
	}
	if org == nil {
		return nil, nil
	}

	domain := "https://" + viper.GetString(config.ServerHostname) +
		":" + viper.GetString(config.ServerPort)

	return &Actor{
		Subject: address,
		Links: []Link{
			{
				Rel:  "self",
				Type: "application/activity+json",
				HREF: domain + "/api/org/" + slug,
			},
		},
	}, nil
}
