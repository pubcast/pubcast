package activitypub

import (
	"context"
	"crypto"
	"net/http"
	"net/url"

	"github.com/go-fed/httpsig"
)

type socialAPI struct{}

func (s socialAPI) ActorIRI(c context.Context, r *http.Request) (*url.URL, error) {
	panic("not implemented")
}

func (s socialAPI) GetSocialAPIVerifier(c context.Context) socialAPIVerifier {
	panic("not implemented")
}

func (s socialAPI) GetPublicKeyForOutbox(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
	panic("not implemented")
}
