package marshal

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalTime(t *testing.T) {
	type SomeStruct struct {
		T MarshalableTime
	}

	var someStruct SomeStruct
	err := json.Unmarshal([]byte(`{"T": "`+time.Now().Format(time.RFC3339)+`"}`), &someStruct)

	assert.NoError(t, err)
}

func TestMarshalTime(t *testing.T) {
	now := time.Now()
	bytes, err := json.Marshal(MarshalableTime{now})
	assert.NoError(t, err)
	assert.Equal(t, now.Format(time.RFC3339), string(bytes))
}
