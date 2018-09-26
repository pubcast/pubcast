package activitypub

import (
	"context"
	"crypto"
	"net/http"
	"net/url"

	"github.com/go-fed/activity/pub"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
)

type socialFederateApplication struct{}

func (social socialFederateApplication) Owns(c context.Context, id *url.URL) bool {
	panic("not implemented")
}

func (social socialFederateApplication) Get(c context.Context, id *url.URL, rw pub.RWType) (pub.PubObject, error) {
	panic("not implemented")
}

func (social socialFederateApplication) GetAsVerifiedUser(c context.Context, id *url.URL, authdUser *url.URL, rw pub.RWType) (pub.PubObject, error) {
	panic("not implemented")
}

func (social socialFederateApplication) Has(c context.Context, id *url.URL) (bool, error) {
	panic("not implemented")
}

func (social socialFederateApplication) Set(c context.Context, o pub.PubObject) error {
	panic("not implemented")
}

func (social socialFederateApplication) GetInbox(c context.Context, r *http.Request, rw pub.RWType) (vocab.OrderedCollectionType, error) {
	panic("not implemented")
}

func (social socialFederateApplication) GetOutbox(c context.Context, r *http.Request, rw pub.RWType) (vocab.OrderedCollectionType, error) {
	panic("not implemented")
}

func (social socialFederateApplication) NewId(c context.Context, t pub.Typer) *url.URL {
	panic("not implemented")
}

func (social socialFederateApplication) GetPublicKey(c context.Context, publicKeyId string) (pubKey crypto.PublicKey, algo httpsig.Algorithm, user *url.URL, err error) {
	panic("not implemented")
}

func (social socialFederateApplication) CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	panic("not implemented")
}

func (social socialFederateApplication) CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	panic("not implemented")
}

func (social socialFederateApplication) ActorIRI(c context.Context, r *http.Request) (*url.URL, error) {
	panic("not implemented")
}

func (social socialFederateApplication) GetSocialAPIVerifier(c context.Context) pub.SocialAPIVerifier {
	panic("not implemented")
}

func (social socialFederateApplication) GetPublicKeyForOutbox(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
	panic("not implemented")
}

func (social socialFederateApplication) OnFollow(c context.Context, s *streams.Follow) pub.FollowResponse {
	panic("not implemented")
}

func (social socialFederateApplication) Unblocked(c context.Context, actorIRIs []*url.URL) error {
	panic("not implemented")
}

func (social socialFederateApplication) FilterForwarding(c context.Context, activity vocab.ActivityType, iris []*url.URL) ([]*url.URL, error) {
	panic("not implemented")
}

func (social socialFederateApplication) NewSigner() (httpsig.Signer, error) {
	panic("not implemented")
}

func (social socialFederateApplication) PrivateKey(boxIRI *url.URL) (privKey crypto.PrivateKey, pubKeyId string, err error) {
	panic("not implemented")
}
