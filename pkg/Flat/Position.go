package flat

import (
	"fmt"

	"github.com/philbrookes/map-generator/pkg/generics"
)

// NewPosition at supplied x, y co-ordinates
func NewPosition(x, y, _ generics.Coordinate) generics.Position {
	return &Position{x: x, y: y}
}

// Position represent a position on a flat plane
type Position struct {
	x generics.Coordinate
	y generics.Coordinate
}

// X value of this position
func (p *Position) X() generics.Coordinate {
	return p.x
}

// Y value of this position
func (p *Position) Y() generics.Coordinate {
	return p.y
}

// Z value of this position (ALWAYS 0)
func (p *Position) Z() generics.Coordinate {
	return 0
}

// XY values of this position
func (p *Position) XY() (generics.Coordinate, generics.Coordinate) {
	return p.x, p.y
}

func (p *Position) Hash() string {
	return fmt.Sprintf("%d0%d", p.x, p.y)
}
