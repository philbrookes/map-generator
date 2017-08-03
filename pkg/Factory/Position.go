package factory

import (
	"github.com/philbrookes/map-generator/pkg/flat"
	"github.com/philbrookes/map-generator/pkg/generics"
)

// Position returns a PositionFunc for creating new positions in the supplied type
func Position(mapType string) generics.PositionFactory {
	switch mapType {
	case "flat":
		return flat.NewPosition
	}

	return flat.NewPosition
}
