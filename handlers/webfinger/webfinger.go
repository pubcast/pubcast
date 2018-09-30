/*
Implements the webfinger protocol.
https://tools.ietf.org/html/rfc7033

This protocol allows one website to ask
another "Do you have an object with this id"?

Webfinger always lives at `/.well-known/webfinger`
and it takes queries like this:

```
/.well-known/webfinger?resource=acct:joe@example.org
```

Webfinger is required for interop'ing with Mastodon
and possibly other sites.
*/

package webfinger

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func getAddressFromResource(resource string) (string, error) {
	fragments := strings.Split(resource, ":")
	if fragments[0] != "acct" {
		return "", errors.New("resource did not start with 'acct:'")
	}

	return fragments[1], nil
}

// Get returns a webfinger response
func Get(w http.ResponseWriter, r *http.Request) {

	// expect a request like ?resource=acct:joe@example.org
	resource := r.URL.Query().Get("resource")
	if resource == "" {
		http.Error(w, "missing 'resource' query parameter in webfinger", http.StatusBadRequest)
		return
	}

	// acct:foo@dogs.com => foo@dogs.com
	address, err := getAddressFromResource(resource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actor, err := atAddress(address)
	if err != nil {

		// Poor formated errors are user errors (ie: 400)
		if _, ok := err.(*badAddressError); ok {
			http.Error(w, "incorrect address format", http.StatusBadRequest)
			return
		}

		// other errors are just 500s
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	str, err := json.Marshal(actor)

	// This _really_ should not happen.
	if err != nil {
		panic(err)
	}

	w.Write(str)
}
