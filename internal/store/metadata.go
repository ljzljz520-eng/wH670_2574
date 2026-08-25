package store

import "time"

type Metadata struct {
	Schema      int
	CreatedAt   time.Time
	Description string
}

func DefaultMetadata() Metadata {
	return Metadata{Schema: 1, CreatedAt: time.Unix(0, 0).UTC(), Description: "language exam archive"}
}
func (m Metadata) Valid() bool { return m.Schema > 0 && m.Description != "" }
