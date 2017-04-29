package Flat

import (
	"github.com/philbrookes/map-generator/pkg/Generics"
	"github.com/pkg/errors"
	"fmt"
)

//Map is a map with a single Z co-ordinate
type Map struct {
	topLeft     Generics.Position
	bottomRight Generics.Position
	tiles 			map[string]Generics.Tile
	structure   string
}


//NewMap creates a new flat map
func NewMap(topLeft, bottomRight Generics.Position) Generics.Map {
	return &Map{topLeft: topLeft, bottomRight: bottomRight, tiles: map[string]Generics.Tile{}, structure: "flat"}
}

// TopLeft returns the top-left position of this flat map
func (m *Map) TopLeft() Generics.Position {
	return m.topLeft
}

// BottomRight returns the bottom-right position of this flat map
func (m *Map) BottomRight() Generics.Position {
	return m.bottomRight
}

func (m *Map) GetTileAt(position Generics.Position) (Generics.Tile, error) {
	if tile, ok := m.tiles[position.Hash()]; ok {
		return tile, nil
	}

	return nil, errors.New(fmt.Sprintf("Could not find tile for position: (%d, %d)\n", position.X(), position.Y()))
}

func (m *Map) IsPositionFilled(position Generics.Position) (bool) {
	if _, err := m.GetTileAt(position); err != nil {
		return false
	}
	return true
}

func(m *Map) AddTile(tile Generics.Tile) error {
	if m.IsPositionFilled(tile.Position()){
		return errors.New(fmt.Sprintf("Tile already placed at: (%d, %d), use UpdateTile instead\n", tile.Position().X(), tile.Position().Y()))
	}

	m.tiles[tile.Position().Hash()] = tile

	return nil
}

func(m *Map) UpdateTile(tile Generics.Tile) {
	m.tiles[tile.Position().Hash()] = tile
}

func(m *Map) Structure() string {
	return m.structure
}