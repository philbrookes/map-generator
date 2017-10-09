package flat

import (
	"github.com/philbrookes/map-generator/pkg/generics"
)

// NewMetadata creates a new metadata object
func NewMetadata(score float32, original bool) generics.Metadata {
	return &Metadata{Score: score, Original: original}
}

// Metadata represent metadata of a single tile
type Metadata struct {
	Score    float32 `json:"score"`
	Original bool    `json:"original"`
}

//GetScore returns the score of this tile
func (m *Metadata) GetScore() float32 {
	return m.Score
}

//GetOriginal returns whether this was an originating tile or not
func (m *Metadata) GetOriginal() bool {
	return m.Original
}

//SetScore changes the current score of this tile
func (m *Metadata) SetScore(s float32) {
	m.Score = s
}

//SetOriginal changes whether this is an originating tile or not
func (m *Metadata) SetOriginal(b bool) {
	m.Original = b
}
