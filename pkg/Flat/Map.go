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
	tiles       map[string]generics.Tile
	structure   string
}

//NewMap creates a new flat map
func NewMap(topLeft, bottomRight generics.Position) generics.Map {
	return &Map{topLeft: topLeft, bottomRight: bottomRight, tiles: map[string]generics.Tile{}, structure: "flat"}
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
	if tile, ok := m.tiles[position.Hash()]; ok {
		return tile, nil
	}

	return nil, errors.New(fmt.Sprintf("Could not find tile for position: (%d, %d)\n", position.X(), position.Y()))
}

func (m *Map) IsPositionFilled(position generics.Position) bool {
	if _, err := m.GetTileAt(position); err != nil {
		return false
	}
	return true
}

func (m *Map) AddTile(tile generics.Tile) error {
	if m.IsPositionFilled(tile.Position()) {
		return errors.New(fmt.Sprintf("Tile already placed at: (%d, %d), use UpdateTile instead\n", tile.Position().X(), tile.Position().Y()))
	}

	m.tiles[tile.Position().Hash()] = tile

	return nil
}

func (m *Map) UpdateTile(tile generics.Tile) {
	m.tiles[tile.Position().Hash()] = tile
}

func (m *Map) Structure() string {
	return m.structure
}
