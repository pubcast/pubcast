package models

import (
	"database/sql"
	"net/url"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/pubcast/pubcast/data"
	"github.com/stretchr/testify/assert"
)

// Runs before everything
func init() {
	data.SetupTestDB()
}

func TestEmptyQueriesSucceed(t *testing.T) {
	db, err := data.NewTestDB()
	assert.NoError(t, err)
	defer db.Close()

	group, err := GetGroup(db, "no-go")
	assert.Nil(t, group)
	assert.Nil(t, err)

	show, err := GetShow(db, "no-go")
	assert.Nil(t, show)
	assert.Nil(t, err)

	pod, err := GetPodcast(db, "no-go")
	assert.Nil(t, pod)
	assert.Nil(t, err)
}

func TestGetGroup(t *testing.T) {
	db, err := data.NewTestDB()
	assert.NoError(t, err)
	defer db.Close()

	// Populate the db with some dummy data
	query := `
		INSERT INTO groups (slug, name, note)
		VALUES ('dog', 'Corgies', 'I like pups')
	`
	_, err = db.Exec(query)
	assert.Nil(t, err) // Inserts should succeed

	group, err := GetGroup(db, "dog")
	assert.Nil(t, err)
	assert.NotNil(t, group)

	assert.Equal(t, "dog", group.Slug, "Group slug should match")
	assert.Equal(t, "Corgies", group.Name, "Group name should match")
	assert.Equal(t, "I like pups", group.Note, "Group note should match")
}

func TestPutGroup(t *testing.T) {
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer db.Close()

	slug, err := PutGroup(db, "hats and ;DROP TABLES", "<html>oh boy</html>")
	assert.Nil(t, err)

	group, err := GetGroup(db, slug)
	assert.Nil(t, err)

	assert.Equal(t, slug, group.Slug)
	assert.Equal(t, "hats and ;DROP TABLES", group.Name)
	assert.Equal(t, "<html>oh boy</html>", group.Note)
}

func TestGetShow(t *testing.T) {
	db, err := data.NewTestDB()
	assert.NoError(t, err)
	defer db.Close()

	// Populate the db with some dummy data
	query := `
		INSERT INTO shows (slug, name, note)
		VALUES ('kitty', 'Cat', 'I like cats')
	`
	_, err = db.Exec(query)
	assert.Nil(t, err) // Inserts should succeed

	show, err := GetShow(db, "kitty")
	assert.Nil(t, err)
	assert.NotNil(t, show)

	assert.Equal(t, "kitty", show.Slug, "Show slug should match")
	assert.Equal(t, "Cat", show.Name, "Show name should match")
	assert.Equal(t, "I like cats", show.Note, "Show note should match")
}

func TestPutShow(t *testing.T) {
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer db.Close()

	slug, err := PutShow(db, "hats and ;DROP TABLES", "<html>oh boy</html>")
	assert.Nil(t, err)

	show, err := GetShow(db, slug)
	assert.Nil(t, err)

	assert.Equal(t, slug, show.Slug)
	assert.Equal(t, "hats and ;DROP TABLES", show.Name)
	assert.Equal(t, "<html>oh boy</html>", show.Note)
}

func sameDay(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func TestGetPodcast(t *testing.T) {
	db, err := data.NewTestDB()
	assert.NoError(t, err)
	defer db.Close()

	// Populate the db with some dummy data
	query := `
		INSERT INTO podcasts (
			slug,
			name,
			note,
			thumbnail_url,
			audio_url,
			media_type,
			posted_at
		) values (
			'foobang',
			'FooBang',
			'some note',
			'https://foo.com/lang.png',
			'https://audio.com/audio.mp3',
			'mp3',
			$1
		)
	`

	now := time.Now()

	_, err = db.Exec(query, now)
	assert.NoError(t, err) // Inserts should succeed

	pod, err := GetPodcast(db, "foobang")
	assert.NoError(t, err)
	assert.Equal(t, "foobang", pod.Slug)
	assert.Equal(t, "FooBang", pod.Name)
	assert.Equal(t, "some note", pod.Note)
	assert.Equal(t, "https://foo.com/lang.png", pod.ThumbnailURL)
	assert.Equal(t, "https://audio.com/audio.mp3", pod.AudioURL)
	assert.Equal(t, mediaType("mp3"), pod.MediaType)
	assert.True(t, sameDay(now.UTC(), pod.PostedAt))
}

func TestPutPodcast(t *testing.T) {
	db, err := data.NewTestDB()
	assert.NoError(t, err)
	defer db.Close()

	thumb, _ := url.Parse("https://thu.mb/nail.png")
	audio, _ := url.Parse("https://audio.com/foo.mp3")

	// Throw something in our db
	slug, err := PutPodcast(db, "name game", "note", thumb, audio, mediaType("mp3"))
	assert.NoError(t, err)
	assert.Equal(t, "name-game", slug)

	podcast, err := GetPodcast(db, slug)
	assert.NoError(t, err)
	assert.Equal(t, "name game", podcast.Name)
	assert.Equal(t, "note", podcast.Note)
	assert.Equal(t, thumb.String(), podcast.ThumbnailURL)
	assert.Equal(t, audio.String(), podcast.AudioURL)
	assert.Equal(t, mediaType("mp3"), podcast.MediaType)
}
