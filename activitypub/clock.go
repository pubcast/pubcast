package activitypub

import "time"

// Implements pub.Clock
type clock struct{}

func (c clock) Now() {
	time.Now()
}
