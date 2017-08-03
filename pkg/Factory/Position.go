package Factory

import (
	"github.com/philbrookes/map-generator/pkg/Flat"
	"github.com/philbrookes/map-generator/pkg/Generics"
)

// Position returns a PositionFunc for creating new positions in the supplied type
func Position(mapType string) Generics.PositionFactory {
	switch mapType {
	case "flat":
		return Flat.NewPosition
	}

	return Flat.NewPosition
}
