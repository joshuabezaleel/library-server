package util

import (
	"github.com/segmentio/ksuid"
)

// NewID returns a string for an ID that can be used for
// all domain models using Segment's ksuid.
func NewID() string {
	return ksuid.New().String()
}
