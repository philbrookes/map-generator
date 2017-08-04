package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/philbrookes/map-generator/pkg/config"
	"github.com/philbrookes/map-generator/pkg/factory"
	"github.com/philbrookes/map-generator/pkg/generics"
	"github.com/pkg/errors"
)

func main() {
	mapType := "flat"
	generatorFactory := factory.Generator(mapType)
	positionFactory := factory.Position(mapType)
	tileFactory := factory.Tile(mapType)
	mapFactory := factory.Map(mapType)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "You must supply an x and y co-ordinate: ./generator <x> <y>\n")
		return
	}

	midX, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(errors.Wrap(err, "X must be numeric"))
	}
	midY, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(errors.Wrap(err, "y must be numeric"))
	}

	buffer := config.GetConfig().GetBufferSize()
	viewport := config.GetConfig().GetViewportSize()
	mapsize := buffer + viewport

	topLeft := positionFactory(generics.Coordinate(midX-(mapsize)), generics.Coordinate(midY-(mapsize)), 0)
	bottomRight := positionFactory(generics.Coordinate(midX+(mapsize)), generics.Coordinate(midY+(mapsize)), 0)

	generator := generatorFactory(topLeft, bottomRight, positionFactory, tileFactory, mapFactory)

	genMap, err := generate(generator)
	if err != nil {
		panic(err)
	}

	err = printMap(genMap, buffer, os.Stdout)
	if err != nil {
		panic(err)
	}

}

func printMap(genMap generics.Map, buffer int, w io.Writer) error {
	bufferCoord := generics.Coordinate(buffer)
	for x := genMap.TopLeft().GetX() + bufferCoord; x <= genMap.BottomRight().GetX()-bufferCoord; x++ {
		for y := genMap.TopLeft().GetY() + bufferCoord; y <= genMap.BottomRight().GetY()-bufferCoord; y++ {
			pos := factory.Position(genMap.Structure())(x, y, 0)
			tile, err := genMap.GetTileAt(pos)
			if err != nil {
				return errors.Wrap(err, "error getting tile")
			}
			switch {
			case tile.GetDanger().GetScore() < 25:
				fmt.Fprint(w, "[#]")
			default:
				fmt.Fprint(w, "[ ]")
			}
		}
		fmt.Fprintln(w, "")
	}
	return nil
}

func generate(generator generics.Generator) (generics.Map, error) {
	genMap, err := generator.Generate()
	if err != nil {
		return nil, errors.Wrap(err, "error generating map")
	}
	return genMap, nil
}
