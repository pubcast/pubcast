**NOTE**:

~This project has been more/less abandoned. I didn't have the time I wanted to commit to this. If you're interested in this project, you're free to reuse the code, the name and the concept.~

I have more time now and am partially reviving this project.

---

# 🎙 Pubcast

[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/metapodcasts)

An _experimental_ (Read: not-usable or in anyway done) distributed/federated podcasting platform based on [ActivityPub.](https://raw.githubusercontent.com/w3c/activitypub/gh-pages/activitypub-tutorial.txt)

## Usage

Getting started:

Ensure that you're using go11 with go-modules turned on.

```sh
export GO111MODULE=on # Put this in your .zshrc or .bash_profile or whatnot
```

Clone/Download the project with:

```sh
go get -u github.com/pubcast/pubcast
```

Building a binary with make (or [mmake](https://github.com/tj/mmake) if you're fancy):

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

Setting up your database (this works best if you have postgres already [running locally](https://postgresapp.com/)):

```sh
make database
```

Creating a new migration in `db/migrations`:

```sh
make migration NAME=some_name_here
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

```json
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

#### Objects, Actors, and Activities

ActivityPub is based on [a formalized vocabulary](https://www.w3.org/TR/activitystreams-vocabulary/) of data types, actions and folks doing the actions.

An `Object` is a generic data type written in JSON:

```json
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Object",
  "id": "http://www.test.example/object/1",
  "name": "A Simple, non-specific object"
}
```

Objects have [a set collection of formalized properties](https://www.w3.org/TR/activitystreams-vocabulary/#properties) such as `id`, `name`, `url`, etc but you technically can create your own. Objects serve as a base type for other Activity Steam's core set of types.

For example, there are a set of [actor types](https://www.w3.org/TR/activitystreams-vocabulary/#actor-types) that themselves are `Objects`.

```json
/* A "Person" actor type */
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Person",
  "name": "Sally Smith"
}
```

[Activities](https://www.w3.org/TR/activitystreams-vocabulary/#h-activity-types) are also subtypes of `Object`, and are used to describe relationships between objects. Some examples of activities include:

- Accept
- Create
- Move
- Question
- Undo
- Follow
- View

An `Activity` json might look something like this:

```json
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally created a note",
  "type": "Create",
  "actor": {
    "type": "Person",
    "name": "Sally"
  },
  "object": {
    "type": "Note",
    "name": "A Simple Note",
    "content": "This is a simple note"
  }
}
```

### Links

- [ActivityPub tutorial](https://raw.githubusercontent.com/w3c/activitypub/gh-pages/activitypub-tutorial.txt)
- [ActivityPub.rocks explaination](https://activitypub.rocks/)
- [W3 ActivityPub Spec](https://www.w3.org/TR/activitypub/)
- [W3 ActivityPub Vocabulary Spec](https://www.w3.org/TR/activitystreams-vocabulary/)
- [Other Golang implementations of this spec](https://github.com/go-fed/activity#who-is-using-this-library-currently)

## Design

### Podcast Object

Podcasts include four main pieces of information: the `header` info, the `shownotes`, the `preview`, and the `audio`. A Header includes the title and date of the show. Shownotes are a collection of info about the show; they're basically an HTML supported description. A preview is an image thumbnail for the show. Audio is the actual stuff you're listening to.

A Podcast ActivityStream Object can therefore look something like this:

```json
"object" : {
 "id": "https://example.org/activity/organization/npr/planet-money",
 "type": "Podcast",
 "name": "This American Life",
 "date": "2008-09-15T15:53:00",
 "shownotes": "Check out our <a href='foo.com'>website!</a>",
 "preview": {
    "type": "Image",
    "href": "http://example.org/album/máiréad.jpg",
    "mediaType": "image/jpeg"
  },
  "audio": {
    "type": "Audio",
    "href": "https://example.org/activity/organization/npr/planet-money/episodes/1.mp4",
    "mediaType": "audio/mp4"
  }
}
```
