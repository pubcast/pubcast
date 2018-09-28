package models

import "github.com/metapods/metapods/data/marshal"

// Group is a collection of Organizations
// Refers to the https://www.w3.org/TR/activitystreams-vocabulary/#dfn-group
// Also refers to the Groups table in the database
type Group struct {
	ID        int                     `json:"id"`
	Slug      string                  `json:"slug"`
	Name      string                  `json:"name"`
	Note      string                  `json:"note"`
	CreatedAt marshal.MarshalableTime `json:"created_at"`
	UpdatedAt marshal.MarshalableTime `json:"updated_at"`
}

// Organization is someone who owns a episodes of podcasts
// Refers to the https://www.w3.org/TR/activitystreams-vocabulary/#dfn-organization
// Also refers to the Organizations table in the database
type Organization struct {
	ID        int                     `json:"id"`
	Slug      string                  `json:"slug"`
	Name      string                  `json:"name"`
	Note      string                  `json:"note"`
	CreatedAt marshal.MarshalableTime `json:"created_at"`
	UpdatedAt marshal.MarshalableTime `json:"updated_at"`
}

// Podcast is a something with an audio link, a name, and a note
// Refers to the Podcasts table in the database
type Podcast struct {
	ID           int                     `json:"id"`
	Slug         string                  `json:"slug"`
	Name         string                  `json:"name"`
	Note         string                  `json:"note"`
	ThumbnailURL string                  `json:"thumbnail_url"`
	AudioURL     string                  `json:"audio_url"`
	MediaType    string                  `json:"media_type"`
	PostedAt     marshal.MarshalableTime `json:"posted_at"`
	CreatedAt    marshal.MarshalableTime `json:"created_at"`
	UpdatedAt    marshal.MarshalableTime `json:"updated_at"`
}
