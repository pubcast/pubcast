package activitypub

import "net/url"

type deliverer struct{}

func (d deliverer) Do(b []byte, to *url.URL, toDo func(b []byte, u *url.URL) error) {
	panic("not implemented")
}
