# Metapod

An _experimental_ (Read: not-usable or in anyway done) distributed podcasting platform based on [ActivityPub.](https://raw.githubusercontent.com/w3c/activitypub/gh-pages/activitypub-tutorial.txt)

## Usage

Getting started:

Install the vendor dependencies using godep. (https://github.com/golang/dep)

```sh
dep ensure
```

Building a binary:

```sh
make
```

Building and then running that binary:

```sh
make run
```

Running tests:

```sh
make test
```

## Learning about ActivityPub

![explaination](https://i.imgur.com/ShgecWe.png)

### Basic Description

ActivityPub gives every user (or `actor` in it's vocab) on a server an "inbox" and an "outbox". But these are really just endpoints:

```
https://myactpub.site/activity/user/flaque/inbox
https://myactpub.site/activity/user/flaque/outbox
```

ActivityPub asks that you accept `GET` and `POST` requests to these endpoints where a `POST` tells a the server to put that in a user's queue or feed and `GET` lets the user retrieve info from the feed. 

You send messages called `ActivityStreams` that are really just a special spec of JSON:

```
{"@context": "https://www.w3.org/ns/activitystreams",
 "type": "Create",
 "id": "https://social.example/alyssa/posts/a29a6843-9feb-4c74-a7f7-081b9c9201d3",
 "to": ["https://chatty.example/ben/"],
 "author": "https://social.example/alyssa/",
 "object": {"type": "Note",
            "id": "https://social.example/alyssa/posts/49e2d03d-b53a-4c4c-a95c-94a6abf45a19",
            "attributedTo": "https://social.example/alyssa/",
            "to": ["https://chatty.example/ben/"],
            "content": "Say, did you finish reading that book I lent you?"}
```


### Links
- [ActivityPub tutorial](https://raw.githubusercontent.com/w3c/activitypub/gh-pages/activitypub-tutorial.txt)
- [ActivityPub.rocks explaination](https://activitypub.rocks/)
- [W3 ActivityPub Spec](https://www.w3.org/TR/activitypub/)
- [Other Golang implementations of this spec](https://github.com/go-fed/activity#who-is-using-this-library-currently)
