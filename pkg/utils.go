package pkg

import (
	"time"

	"github.com/bouk/monkey"
	"github.com/segmentio/ksuid"
)

// NewID returns a string for an ID that can be used for
// all domain models using Segment's ksuid.
func NewID() string {
	return ksuid.New().String()
}

// CreatedTimePatch returns the patched time for creating a
// timestamp.
func CreatedTimePatch() (time.Time, *monkey.PatchGuard) {
	createdTime := time.Now()
	timePatch := monkey.Patch(time.Now, func() time.Time {
		return createdTime
	})

	return createdTime, timePatch
}

// NewIDPatch returns the patched ID for creating a
// ID.
func NewIDPatch() (string, *monkey.PatchGuard) {
	ID := NewID()
	IDPatch := monkey.Patch(NewID, func() string {
		return ID
	})

	return ID, IDPatch
}
