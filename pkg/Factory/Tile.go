package Factory

import (
"github.com/philbrookes/map-generator/pkg/Flat"
"github.com/philbrookes/map-generator/pkg/Generics"
)

// Tile returns a TileFunc for creating new Tiles in the supplied type
func Tile(mapType string) Generics.TileFactory {
	switch mapType {
	case "flat":
		return Flat.NewTile
	}

	return Flat.NewTile
}

