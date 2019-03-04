# What's Webfinger?

[Webfinger](https://en.wikipedia.org/wiki/WebFinger), a protocol for discovering objects on the server. It's used by Mastodon and is important for interop'ing with Mastodon (and most ActivityPub servers).

It lives at a special route: `GET /.well-known/webfinger`.

We can expect a webfinger response to always looks something like this:

```json
{
  "subject": "acct:alice@my-example.com",

  "links": [
    {
      "rel": "self",
      "type": "application/activity+json",
      "href": "https://my-example.com/actor"
    }
  ]
}
```

In this case, `alice` is the ActivityPub Organization slug, and `my-example.com` is the domain of the server.
