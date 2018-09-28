package marshal

import (
	"strings"
	"time"
)

// MarshalableTime is an extension to time that's JSON marshalable
type MarshalableTime time.Time

// UnmarshalJSON marshals time in RFC3339 format to json
func (m *MarshalableTime) UnmarshalJSON(p []byte) error {
	t, err := time.Parse(time.RFC3339, strings.Replace(
		string(p),
		"\"",
		"",
		-1,
	))

	if err != nil {
		return err
	}

	*m = MarshalableTime(t)

	return nil
}
