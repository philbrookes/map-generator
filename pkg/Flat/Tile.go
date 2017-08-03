package flat

import (
	"github.com/philbrookes/map-generator/pkg/generics"
)

// NewTile at supplied position
func NewTile(position generics.Position) generics.Tile {
	metadata := NewMetadata(0, false)
	return &Tile{position: position, danger: metadata}
}

type Tile struct {
	position generics.Position
	danger   generics.Metadata
}

func (t *Tile) Position() generics.Position {
	return t.position
}

func (t *Tile) Danger() generics.Metadata {
	return t.danger
}
