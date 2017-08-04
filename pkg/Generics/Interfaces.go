package generics

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
	GetPosition() Position
	GetDanger() Metadata
}

//Metadata is an interface defining the metadata of a tile
type Metadata interface {
	GetScore() float32
	GetOriginal() bool
	SetScore(float32)
	SetOriginal(bool)
}

// Position is an interface to a map position
type Position interface {
	GetX() Coordinate
	GetY() Coordinate
	GetZ() Coordinate
	GetXY() (Coordinate, Coordinate)
	GetHash() string
}

// Coordinate represents a value on a dimension
type Coordinate int
