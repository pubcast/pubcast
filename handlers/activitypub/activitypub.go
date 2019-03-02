package activitypub

import (
	"context"
	"net/http"

	"github.com/go-fed/activity/streams/vocab"
)

// onCreate implements the go-fed activity/streams/vocab JSON Resolver
func onCreate(c context.Context, create vocab.ActivityStreamsCreate) error {
	return nil
}

// onUpdate implements the go-fed activity/streams/vocab JSON Resolver
func onUpdate(c context.Context, update vocab.ActivityStreamsUpdate) error {
	return nil
}

// Outbox implements the ActivityPub actor outbox
// See: https://www.w3.org/TR/activitypub/#outbox
func Outbox(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unimplemented", http.StatusNotImplemented)
	return
}

// Inbox implements the ActivityPub actor inbox
// See: https://www.w3.org/TR/activitypub/#inbox
func Inbox(w http.ResponseWriter, r *http.Request) {

	http.Error(w, "Unimplemented", http.StatusNotImplemented)
	return

	// TODO GET TO THIS
	//
	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }

	// jsonResolver, err := streams.NewJSONResolver(onCreate, onUpdate)
	// if err != nil {
	// 	// Something in the setup was wrong. For example, a callback has an
	// 	// unsupported signature and would never be called
	// 	http.Error(w, "JSON Resolver not setup correctly; contact a developer immediately", http.StatusInternalServerError)
	// }

	// var jsonMap map[string]interface{}
	// if err = json.Unmarshal(b, &jsonMap); err != nil {
	// 	panic(err) // TODO don't panic here
	// }

	// c := context.Background()
	// err = jsonResolver.Resolve(c, jsonMap)
	// if err != nil && !streams.IsUnmatchedErr(err) {
	// 	// Something went wrong
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// } else if streams.IsUnmatchedErr(err) {
	// 	// Everything went right but the callback didn't match or the ActivityStreams
	// 	// type is one that wasn't code generated.
	// 	http.Error(w, "not an activityStreams object", http.StatusBadRequest)
	// }
}
