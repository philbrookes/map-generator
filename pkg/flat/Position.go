package flat

import (
	"fmt"

	"github.com/philbrookes/map-generator/pkg/generics"
)

// NewPosition at supplied x, y co-ordinates
func NewPosition(x, y, _ generics.Coordinate) generics.Position {
	return &Position{X: x, Y: y}
}

// Position represent a position on a flat plane
type Position struct {
	X generics.Coordinate `json:"x"`
	Y generics.Coordinate `json:"y"`
}

// GetX value of this position
func (p *Position) GetX() generics.Coordinate {
	return p.X
}

// GetY value of this position
func (p *Position) GetY() generics.Coordinate {
	return p.Y
}

// GetZ value of this position (ALWAYS 0)
func (p *Position) GetZ() generics.Coordinate {
	return 0
}

// GetXY values of this position
func (p *Position) GetXY() (generics.Coordinate, generics.Coordinate) {
	return p.X, p.Y
}

//GetHash returns hash of this position
func (p *Position) GetHash() string {
	return fmt.Sprintf("%d0%d", p.X, p.Y)

}
