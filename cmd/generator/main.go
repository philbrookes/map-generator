package main

import (
	"fmt"
	"github.com/philbrookes/map-generator/pkg/Factory"
	"github.com/philbrookes/map-generator/pkg/Generics"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
)

func main() {
	mapType := "flat"
	generatorFactory := Factory.Generator(mapType)
	positionFactory := Factory.Position(mapType)
	tileFactory := Factory.Tile(mapType)
	mapFactory := Factory.Map(mapType)

	midX, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(errors.Wrap(err, "X must be numeric"))
	}
	midY, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(errors.Wrap(err, "y must be numeric"))
	}

	buffer := 15
	viewport := 25
	mapsize := buffer + viewport

	topLeft := positionFactory(Generics.Coordinate(midX-(mapsize)), Generics.Coordinate(midY-(mapsize)), 0)
	bottomRight := positionFactory(Generics.Coordinate(midX+(mapsize)), Generics.Coordinate(midY+(mapsize)), 0)

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

func printMap(genMap Generics.Map, buffer int, w io.Writer) error {
	bufferCoord := Generics.Coordinate(buffer)
	for x:=genMap.TopLeft().X()+bufferCoord;x<=genMap.BottomRight().X()-bufferCoord;x++{
		for y:=genMap.TopLeft().Y()+bufferCoord;y<=genMap.BottomRight().Y()-bufferCoord;y++{
			pos := Factory.Position(genMap.Structure())(x, y, 0)
			tile, err := genMap.GetTileAt(pos)
			if err != nil {
				return errors.Wrap(err, "error getting tile")
			}
			switch {
			case tile.Danger().Score() < 25:
				fmt.Fprint(w, "[#]")
			default:
				fmt.Fprint(w, "[ ]")
			}
		}
		fmt.Fprintln(w, "")
	}
	return nil
}


func generate(generator Generics.Generator) (Generics.Map, error) {
	genMap, err := generator.Generate()
	if err != nil {
		return nil, errors.Wrap(err, "error generating map")
	}
	return genMap, nil
}
