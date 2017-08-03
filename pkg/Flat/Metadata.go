package Flat

import (
	"github.com/philbrookes/map-generator/pkg/Generics"
)

// NewMetadata creates a new metadata object
func NewMetadata(score float32, original bool) Generics.Metadata {
	return &Metadata{score: score, original: original}
}

// Metadata represent metadata of a single tile
type Metadata struct {
	score float32
	original bool
}

func (m *Metadata) Score() float32 {
	return m.score
}

func (m *Metadata) Original() bool {
	return m.original
}

func (m *Metadata) SetScore(s float32) {
	m.score = s
}

func (m *Metadata) SetOriginal(b bool) {
	m.original = b
}