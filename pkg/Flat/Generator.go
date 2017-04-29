package Flat

import (
	"github.com/philbrookes/map-generator/pkg/Generics"
	"math/rand"
	"github.com/pkg/errors"
	"fmt"
	"strconv"
	"math"
)

//NewGenerator creates a new flat map generator
func NewGenerator(topLeft, bottomRight Generics.Position, positionFactory Generics.PositionFactory, tileFactory Generics.TileFactory, mapFactory Generics.MapFactory) Generics.Generator {
	seed, _ := strconv.Atoi(fmt.Sprintf("%d%d%d%d", topLeft.X(), topLeft.Y(), bottomRight.X(), bottomRight.Y()))
	rng := rand.New(rand.NewSource(int64(seed)))
	return &Generator{topLeft: topLeft, bottomRight: bottomRight, positionFactory: positionFactory, tileFactory: tileFactory, mapFactory: mapFactory, rng: rng}
}

// Generator generates maps
type Generator struct {
	topLeft Generics.Position
	bottomRight Generics.Position
	tileFactory Generics.TileFactory
	positionFactory Generics.PositionFactory
	mapFactory Generics.MapFactory
	rng *rand.Rand
}

// Generate a Flat.Map
func (g *Generator) Generate() (Generics.Map, error) {
	genMap := g.mapFactory(g.topLeft, g.bottomRight)

	for x:=g.topLeft.X();x<=g.bottomRight.X();x++{
		for y:=g.topLeft.Y();y<=g.bottomRight.Y();y++{
			tile := g.tileFactory(g.positionFactory(x, y, 0))
			err := genMap.AddTile(tile)
			if err != nil {
				return nil, errors.Wrap(err, "Problem adding tile to map")
			}
			hash, _ := strconv.Atoi(tile.Position().Hash())
			g.rng.Seed(int64(hash))
			content := g.rng.Int() % 40
			if content == 0 {
				tile.Danger().SetOriginal(true)
				tile.Danger().SetScore(100)
			}
		}
	}

	warmedMap, err := g.warm(genMap, 15)
	if err != nil {
		return nil, errors.Wrap(err, "error warming map")
	}
	return warmedMap, nil
}


//recursive
func(g *Generator) warm(spotMap Generics.Map, steps int) (Generics.Map, error) {
	if steps <= 0 {
		return spotMap, nil
	}

	warmedMap := g.mapFactory(g.topLeft, g.bottomRight)
	for x:=spotMap.TopLeft().X();x<=spotMap.BottomRight().X();x++ {
		for y := spotMap.TopLeft().Y(); y <= spotMap.BottomRight().Y(); y++ {
			total := float32(0)
			tileCount := 0
			startX := Generics.Coordinate(math.Max(float64(x-1), float64(spotMap.TopLeft().X())))
			startY := Generics.Coordinate(math.Max(float64(y-1), float64(spotMap.TopLeft().Y())))
			stopX := Generics.Coordinate(math.Min(float64(x+1), float64(spotMap.BottomRight().X())))
			stopY := Generics.Coordinate(math.Min(float64(y+1), float64(spotMap.BottomRight().Y())))

			tile, err := spotMap.GetTileAt(g.positionFactory(x, y, 0))
			if err != nil {
				return nil, errors.Wrap(err, "Could not load tile")
			}
			if tile.Danger().Original() {
				warmedMap.AddTile(tile)
				continue
			}

			for scanX:=startX;scanX<=stopX;scanX++{
				for scanY:=startY;scanY<=stopY;scanY++{
					scanTile, err := spotMap.GetTileAt(g.positionFactory(scanX, scanY, 0))
					if err != nil {
						return nil, errors.Wrap(err, "error loading scan tile")
					}
					total += scanTile.Danger().Score()
					tileCount++
				}
			}
			tile.Danger().SetScore(total / float32(tileCount))
			err = warmedMap.AddTile(tile)
			if err != nil {
				return nil, errors.Wrap(err, "Problem adding tile to warmed map")
			}
		}
	}

	return g.warm(warmedMap, steps-1)
}
