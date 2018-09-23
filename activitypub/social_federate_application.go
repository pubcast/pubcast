import (
	"context"
	"crypto"
	"net/http"
	"net/url"

	"github.com/go-fed/activity/pub"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
)

func (s socialFederateApplication) Owns(c context.Context, id *url.URL) bool {
	panic("not implemented")
}

func (s socialFederateApplication) Get(c context.Context, id *url.URL, rw pub.RWType) (pub.PubObject, error) {
	panic("not implemented")
}

func (s socialFederateApplication) GetAsVerifiedUser(c context.Context, id *url.URL, authdUser *url.URL, rw pub.RWType) (pub.PubObject, error) {
	panic("not implemented")
}

func (s socialFederateApplication) Has(c context.Context, id *url.URL) (bool, error) {
	panic("not implemented")
}

func (s socialFederateApplication) Set(c context.Context, o pub.PubObject) error {
	panic("not implemented")
}

func (s socialFederateApplication) GetInbox(c context.Context, r *http.Request, rw pub.RWType) (vocab.OrderedCollectionType, error) {
	panic("not implemented")
}

func (s socialFederateApplication) GetOutbox(c context.Context, r *http.Request, rw pub.RWType) (vocab.OrderedCollectionType, error) {
	panic("not implemented")
}

func (s socialFederateApplication) NewId(c context.Context, t pub.Typer) *url.URL {
	panic("not implemented")
}

func (s socialFederateApplication) GetPublicKey(c context.Context, publicKeyId string) (pubKey crypto.PublicKey, algo httpsig.Algorithm, user *url.URL, err error) {
	panic("not implemented")
}

func (s socialFederateApplication) CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	panic("not implemented")
}

func (s socialFederateApplication) CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	panic("not implemented")
}

func (s socialFederateApplication) ActorIRI(c context.Context, r *http.Request) (*url.URL, error) {
	panic("not implemented")
}

func (s socialFederateApplication) GetSocialAPIVerifier(c context.Context) pub.SocialAPIVerifier {
	panic("not implemented")
}

func (s socialFederateApplication) GetPublicKeyForOutbox(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
	panic("not implemented")
}

func (s socialFederateApplication) OnFollow(c context.Context, s *streams.Follow) pub.FollowResponse {
	panic("not implemented")
}

func (s socialFederateApplication) Unblocked(c context.Context, actorIRIs []*url.URL) error {
	panic("not implemented")
}

func (s socialFederateApplication) FilterForwarding(c context.Context, activity vocab.ActivityType, iris []*url.URL) ([]*url.URL, error) {
	panic("not implemented")
}

func (s socialFederateApplication) NewSigner() (httpsig.Signer, error) {
	panic("not implemented")
}

func (s socialFederateApplication) PrivateKey(boxIRI *url.URL) (privKey crypto.PrivateKey, pubKeyId string, err error) {
	panic("not implemented")
}

