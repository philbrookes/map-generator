package Generics

// Generator is an interface of a map generator
type Generator interface {
	Generate() (Map, error)
}

// Map is a collection of connected Tiles
type Map interface {
	TopLeft() Position
	BottomRight() Position
	AddTile(tile Tile) error
	GetTileAt(position Position) (Tile, error)
	IsPositionFilled(position Position) bool
	UpdateTile(tile Tile)
	Structure() string
}

// Tile is an interface of a plot on a map
type Tile interface {
	Position() Position
	Danger() Metadata
}

type Metadata interface {
	Score() float32
	Original() bool
	SetScore(float32)
	SetOriginal(bool)
}

// Position is an interface to a map position
type Position interface {
	X() Coordinate
	Y() Coordinate
	Z() Coordinate
	XY() (Coordinate, Coordinate)
	Hash() string
}

// Coordinate represents a value on a dimension
type Coordinate int
