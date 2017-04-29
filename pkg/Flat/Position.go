package Flat

import (
	"github.com/philbrookes/map-generator/pkg/Generics"
	"fmt"
)

// NewPosition at supplied x, y co-ordinates
func NewPosition(x, y, _ Generics.Coordinate) Generics.Position {
	return &Position{x: x, y: y}
}

// Position represent a position on a flat plane
type Position struct {
	x Generics.Coordinate
	y Generics.Coordinate
}

// X value of this position
func (p *Position) X() Generics.Coordinate {
	return p.x
}

// Y value of this position
func (p *Position) Y() Generics.Coordinate {
	return p.y
}

// Z value of this position (ALWAYS 0)
func (p *Position) Z() Generics.Coordinate {
	return 0
}

// XY values of this position
func (p *Position) XY() (Generics.Coordinate, Generics.Coordinate) {
	return p.x, p.y
}

func (p *Position) Hash() string {
	return fmt.Sprintf("%d0%d", p.x, p.y)
}