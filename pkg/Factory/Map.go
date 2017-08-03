package factory

import (
	"github.com/philbrookes/map-generator/pkg/flat"
	"github.com/philbrookes/map-generator/pkg/generics"
)

// Map returns a MapFactory for creating new maps in the supplied type
func Map(mapType string) generics.MapFactory {
	switch mapType {
	case "flat":
		return flat.NewMap
	}

	return flat.NewMap
}

