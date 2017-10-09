package factory

import "github.com/philbrookes/map-generator/pkg/generics"
import "github.com/philbrookes/map-generator/pkg/flat"

// Tile returns a TileFunc for creating new Tiles in the supplied type
func Tile(mapType string) generics.TileFactory {
	switch mapType {
	case "flat":
		return flat.NewTile
	}

	return flat.NewTile
}
