package flat

import (
	"github.com/philbrookes/map-generator/pkg/generics"
)

// NewTile at supplied position
func NewTile(position generics.Position) generics.Tile {
	metadata := NewMetadata(0, false)
	return &Tile{Position: position, Danger: metadata}
}

type Tile struct {
	Position generics.Position `json:"position"`
	Danger   generics.Metadata `json:"danger"`
}

func (t *Tile) GetPosition() generics.Position {
	return t.Position
}

func (t *Tile) GetDanger() generics.Metadata {
	return t.Danger
}
