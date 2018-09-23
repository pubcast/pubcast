package activitypub

import "net/http"

type httpClient struct{}

func (h httpClient) Do(req *http.Request) (*http.Response, error) {
	panic("not implemented")
}
