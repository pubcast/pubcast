package activitypub

import (
	"context"
	"crypto"
	"net/url"

	"github.com/go-fed/activity/pub"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
)

type federateAPI struct{}

func (f federateAPI) OnFollow(c context.Context, s *streams.Follow) pub.FollowResponse {
	panic("not implemented")
}

func (f federateAPI) Unblocked(c context.Context, actorIRIs []*url.URL) error {
	panic("not implemented")
}

func (f federateAPI) FilterForwarding(c context.Context, activity vocab.ActivityType, iris []*url.URL) ([]*url.URL, error) {
	panic("not implemented")
}

func (f federateAPI) NewSigner() (httpsig.Signer, error) {
	panic("not implemented")
}

func (f federateAPI) PrivateKey(boxIRI *url.URL) (privKey crypto.PrivateKey, pubKeyId string, err error) {
	panic("not implemented")
}
