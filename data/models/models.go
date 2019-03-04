package models

import (
	"database/sql"
	"net/url"
	"time"

	slugify "github.com/gosimple/slug"
)

// Group is a collection of Organizations
// Refers to the https://www.w3.org/TR/activitystreams-vocabulary/#dfn-group
// Also refers to the Groups table in the database
type Group struct {
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetGroup returns a single Group object or nil
func GetGroup(db *sql.DB, slug string) (*Group, error) {
	row := db.QueryRow(`
		select slug, name, note, created_at, updated_at
		from groups where slug = $1
	`, slug)

	var group Group
	err := row.Scan(&group.Slug,
		&group.Name, &group.Note, &group.CreatedAt, &group.UpdatedAt)

	// This is not an error from the user's perspective
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// PutGroup creates a group with this name and note
func PutGroup(db *sql.DB, name string, note string) (string, error) {
	slug := slugify.MakeLang(name, "en")

	query := `
		INSERT INTO groups (slug, name, note)
		VALUES ($1, $2, $3)
	`

	_, err := db.Exec(query, slug, name, note)
	return slug, err
}

// A Show is someone who owns a episodes of podcasts
// Refers to the https://www.w3.org/TR/activitystreams-vocabulary/#dfn-organization
// Also refers to the Show table in the database
type Show struct {
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetShow gets an Show at any slug
func GetShow(db *sql.DB, slug string) (*Show, error) {
	row := db.QueryRow(`
		SELECT slug, name, note, created_at, updated_at
		FROM shows WHERE slug = $1
	`, slug)

	var show Show
	err := row.Scan(&show.Slug,
		&show.Name, &show.Note, &show.CreatedAt, &show.UpdatedAt)

	// This is not an error from the user's perspective
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &show, nil
}

// PutShow creates a show with this name and note
func PutShow(db *sql.DB, name string, note string) (string, error) {
	slug := slugify.MakeLang(name, "en")

	query := `
		INSERT INTO shows (slug, name, note)
		VALUES ($1, $2, $3)
	`

	_, err := db.Exec(query, slug, name, note)
	return slug, err
}

// All supported media types
type mediaType string

const (
	// mp3 is recommended
	mp3 mediaType = "mp3"
	m4a mediaType = "m4a"
	ogg mediaType = "ogg"
)

// Podcast is a something with an audio link, a name, and a note
// Refers to the Podcasts table in the database
type Podcast struct {
	Slug         string    `json:"slug"`
	Name         string    `json:"name"`
	Note         string    `json:"note"`
	ThumbnailURL string    `json:"thumbnail_url"`
	AudioURL     string    `json:"audio_url"`
	MediaType    mediaType `json:"media_type"`
	PostedAt     time.Time `json:"posted_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// PutPodcast adds a podcast to the database
// It generates a slug and posted_at timestamp
func PutPodcast(
	db *sql.DB,
	name string,
	note string,
	thumbnailURL *url.URL,
	audioURL *url.URL,
	media mediaType,
) (string, error) {
	slug := slugify.MakeLang(name, "en")

	query := `
		INSERT INTO podcasts (
			slug,
			name,
			note,
			thumbnail_url,
			audio_url,
			media_type,
			posted_at
		) Values ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := db.Exec(query, slug, name, note, thumbnailURL.String(), audioURL.String(), media, time.Now())
	return slug, err
}

// GetPodcast queries a db looking for a podcast
func GetPodcast(db *sql.DB, slug string) (*Podcast, error) {
	row := db.QueryRow(`
		SELECT slug, name, note, thumbnail_url, audio_url, media_type, posted_at, created_at, updated_at FROM podcasts WHERE slug = $1;
	`, slug)

	var pod Podcast
	err := row.Scan(
		&pod.Slug,
		&pod.Name,
		&pod.Note,
		&pod.ThumbnailURL,
		&pod.AudioURL,
		&pod.MediaType,
		&pod.PostedAt,
		&pod.CreatedAt,
		&pod.UpdatedAt,
	)

	// This is not an error from the user's perspective
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &pod, nil
}
