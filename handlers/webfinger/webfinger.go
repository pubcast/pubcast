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

import "net/url"

func Get(id *url.URL) {

}
