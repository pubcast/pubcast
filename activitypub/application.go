package activitypub

import (
	"context"
	"crypto"
	"net/http"
	"net/url"

	"github.com/Flaque/metapod/config"
	"github.com/go-fed/activity/pub"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
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

// Determines if the app has ActivityStream data at the IRI (Internationalized Resource ID)
func (a application) Has(c context.Context, id *url.URL) (bool, error) {
	panic("not implemented")
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
