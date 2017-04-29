package Flat


import (
	"github.com/philbrookes/map-generator/pkg/Generics"
)

// NewTile at supplied position
func NewTile(position Generics.Position) Generics.Tile {
	metadata := NewMetadata(0, false)
	return &Tile{position: position, danger: metadata}
}

type Tile struct {
	position Generics.Position
	danger Generics.Metadata
}

func (t *Tile) Position() Generics.Position {
	return t.position
}

func (t *Tile) Danger() Generics.Metadata {
	return t.danger
}
