package Factory

import (
	"github.com/philbrookes/map-generator/pkg/Flat"
	"github.com/philbrookes/map-generator/pkg/Generics"
)

// Map returns a MapFactory for creating new maps in the supplied type
func Map(mapType string) Generics.MapFactory {
	switch mapType {
	case "flat":
		return Flat.NewMap
	}

	return Flat.NewMap
}

