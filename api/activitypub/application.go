package activitypub

import (
	"context"
	"crypto"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-fed/activity/pub"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
	"github.com/metapods/metapods/config"
	db "github.com/metapods/metapods/data"
	"github.com/metapods/metapods/data/models"
	"github.com/spf13/viper"
)

type application struct{}

// Determines wether the app owns an IRI, or Internationalized Resource ID
func (a application) Owns(c context.Context, id *url.URL) bool {
	actual := id.Host
	expected := viper.GetString(config.ServerHostname) + ":" + viper.GetString(config.ServerPort)

	return actual == expected
}

// Gets ActivityStream content
func (a application) Get(c context.Context, id *url.URL, rw pub.RWType) (pub.PubObject, error) {
	panic("not implemented")
}

func (a application) GetAsVerifiedUser(c context.Context, id *url.URL, authdUser *url.URL, rw pub.RWType) (pub.PubObject, error) {
	panic("not implemented")
}

// Deletes empty strings from an array of strings
// ["", "dogs", "oh"] => ["dogs", "oh"]
func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func matchesURLSpec(id *url.URL) bool {
	fragments := deleteEmpty(strings.Split(id.Path, "/"))

	// Fragments need four pieces "activity", "<object>", "<value>", "inbox or outbox"
	if len(fragments) != 4 {
		fmt.Println("length", len(fragments))
		return false
	}

	// We only accept paths starting with /activity
	if fragments[0] != "activity" {
		return false
	}

	// The ActivityStreams object we're trying to reference
	object := fragments[1]
	if !(object == "organization" || object == "group") {
		return false
	}

	// The forth piece must be "inbox" or "outbox"
	if !(fragments[3] == "inbox" || fragments[3] == "outbox") {
		return false
	}

	return true
}

func getSlug(id *url.URL) string {
	fragments := deleteEmpty(strings.Split(id.Path, "/"))
	if len(fragments) != 4 {
		return "" // Likely bad.
	}

	return fragments[2]
}

// Determines if the app has ActivityStream data at the IRI (Internationalized Resource ID)
// We expect IRIs to have a path like `/activity/<object>/<value>/<inbox|outbox>`
func (a application) Has(c context.Context, id *url.URL) (bool, error) {
	if !matchesURLSpec(id) {
		return false, nil
	}

	group, err := models.GetGroup(db.Pool, getSlug(id))
	if err != nil {
		return false, err
	}
	if group == nil {
		return false, nil
	}

	return true, nil
}

// Sets the ActivityStream data
func (a application) Set(c context.Context, o pub.PubObject) error {
	panic("not implemented")
}

func (a application) GetInbox(c context.Context, r *http.Request, rw pub.RWType) (vocab.OrderedCollectionType, error) {
	panic("not implemented")
}

func (a application) GetOutbox(c context.Context, r *http.Request, rw pub.RWType) (vocab.OrderedCollectionType, error) {
	panic("not implemented")
}

func (a application) NewId(c context.Context, t pub.Typer) *url.URL {
	panic("not implemented")
}

func (a application) GetPublicKey(c context.Context, publicKeyId string) (pubKey crypto.PublicKey, algo httpsig.Algorithm, user *url.URL, err error) {
	panic("not implemented")
}

func (a application) CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	panic("not implemented")
}

func (a application) CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	panic("not implemented")
}
