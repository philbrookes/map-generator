package flat

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/philbrookes/map-generator/pkg/generics"
	"github.com/pkg/errors"
)

//NewGenerator creates a new flat map generator
func NewGenerator(topLeft, bottomRight generics.Position, positionFactory generics.PositionFactory, tileFactory generics.TileFactory, mapFactory generics.MapFactory) generics.Generator {
	seed, _ := strconv.Atoi(fmt.Sprintf("%d%d%d%d", topLeft.GetX(), topLeft.GetY(), bottomRight.GetX(), bottomRight.GetY()))
	rng := rand.New(rand.NewSource(int64(seed)))
	return &Generator{topLeft: topLeft, bottomRight: bottomRight, positionFactory: positionFactory, tileFactory: tileFactory, mapFactory: mapFactory, rng: rng}
}

// Generator generates maps
type Generator struct {
	topLeft         generics.Position
	bottomRight     generics.Position
	tileFactory     generics.TileFactory
	positionFactory generics.PositionFactory
	mapFactory      generics.MapFactory
	rng             *rand.Rand
}

// Generate a flat.Map
func (g *Generator) Generate() (generics.Map, error) {
	genMap := g.mapFactory(g.topLeft, g.bottomRight)

	for x := g.topLeft.GetX(); x <= g.bottomRight.GetX(); x++ {
		for y := g.topLeft.GetY(); y <= g.bottomRight.GetY(); y++ {
			tile := g.tileFactory(g.positionFactory(x, y, 0))
			err := genMap.AddTile(tile)
			if err != nil {
				return nil, errors.Wrap(err, "Problem adding tile to map")
			}
			hash, _ := strconv.Atoi(tile.GetPosition().GetHash())
			g.rng.Seed(int64(hash))
			content := g.rng.Int() % 40
			if content == 0 {
				tile.GetDanger().SetOriginal(true)
				tile.GetDanger().SetScore(100)
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
func (g *Generator) warm(spotMap generics.Map, steps int) (generics.Map, error) {
	if steps <= 0 {
		return spotMap, nil
	}

	warmedMap := g.mapFactory(g.topLeft, g.bottomRight)
	for x := spotMap.TopLeft().GetX(); x <= spotMap.BottomRight().GetX(); x++ {
		for y := spotMap.TopLeft().GetY(); y <= spotMap.BottomRight().GetY(); y++ {
			total := float32(0)
			tileCount := 0
			startX := generics.Coordinate(math.Max(float64(x-1), float64(spotMap.TopLeft().GetX())))
			startY := generics.Coordinate(math.Max(float64(y-1), float64(spotMap.TopLeft().GetY())))
			stopX := generics.Coordinate(math.Min(float64(x+1), float64(spotMap.BottomRight().GetX())))
			stopY := generics.Coordinate(math.Min(float64(y+1), float64(spotMap.BottomRight().GetY())))

			tile, err := spotMap.GetTileAt(g.positionFactory(x, y, 0))
			if err != nil {
				return nil, errors.Wrap(err, "Could not load tile")
			}
			if tile.GetDanger().GetOriginal() {
				warmedMap.AddTile(tile)
				continue
			}

			for scanX := startX; scanX <= stopX; scanX++ {
				for scanY := startY; scanY <= stopY; scanY++ {
					scanTile, err := spotMap.GetTileAt(g.positionFactory(scanX, scanY, 0))
					if err != nil {
						return nil, errors.Wrap(err, "error loading scan tile")
					}
					total += scanTile.GetDanger().GetScore()
					tileCount++
				}
			}
			tile.GetDanger().SetScore(total / float32(tileCount))
			err = warmedMap.AddTile(tile)
			if err != nil {
				return nil, errors.Wrap(err, "Problem adding tile to warmed map")
			}
		}
	}

	return g.warm(warmedMap, steps-1)
}
