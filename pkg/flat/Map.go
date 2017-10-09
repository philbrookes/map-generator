package flat

import (
	"fmt"

	"github.com/philbrookes/map-generator/pkg/generics"
	"github.com/pkg/errors"
)

//Map is a map with a single Z co-ordinate
type Map struct {
	topLeft     generics.Position
	bottomRight generics.Position
	Tiles       map[string]generics.Tile `json:"tiles"`
	structure   string
}

//NewMap creates a new flat map
func NewMap(topLeft, bottomRight generics.Position) generics.Map {
	return &Map{topLeft: topLeft, bottomRight: bottomRight, Tiles: map[string]generics.Tile{}, structure: "flat"}
}

// TopLeft returns the top-left position of this flat map
func (m *Map) TopLeft() generics.Position {
	return m.topLeft
}

// BottomRight returns the bottom-right position of this flat map
func (m *Map) BottomRight() generics.Position {
	return m.bottomRight
}

func (m *Map) GetTileAt(position generics.Position) (generics.Tile, error) {
	if tile, ok := m.Tiles[position.GetHash()]; ok {
		return tile, nil
	}

	return nil, errors.New(fmt.Sprintf("Could not find tile for position: (%d, %d)\n", position.GetX(), position.GetY()))
}

func (m *Map) IsPositionFilled(position generics.Position) bool {
	if _, err := m.GetTileAt(position); err != nil {
		return false
	}
	return true
}

func (m *Map) AddTile(tile generics.Tile) error {
	if m.IsPositionFilled(tile.GetPosition()) {
		return errors.New(fmt.Sprintf("Tile already placed at: (%d, %d), use UpdateTile instead\n", tile.GetPosition().GetX(), tile.GetPosition().GetY()))
	}

	m.Tiles[tile.GetPosition().GetHash()] = tile

	return nil
}

func (m *Map) UpdateTile(tile generics.Tile) {
	m.Tiles[tile.GetPosition().GetHash()] = tile
}

func (m *Map) Structure() string {
	return m.structure
}
