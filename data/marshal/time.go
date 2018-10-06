package marshal

import (
	"time"
)

// MarshalableTime is an extension to time that's JSON marshalable
type MarshalableTime struct {
	time.Time
}

// UnmarshalJSON converts json time in RFC3339 format to a go object
func (m *MarshalableTime) UnmarshalJSON(p []byte) error {
	t, err := time.Parse(time.RFC3339, p)

	if err != nil {
		return err
	}

	*m = MarshalableTime{t}
	return nil
}

// MarshalJSON turns a go time object into a json string
func (m *MarshalableTime) MarshalJSON() ([]byte, error) {
	return []byte(m.Format(time.RFC3339)), nil
}
