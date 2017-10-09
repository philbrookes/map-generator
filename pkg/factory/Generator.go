package factory

import (
	"github.com/philbrookes/map-generator/pkg/flat"
	"github.com/philbrookes/map-generator/pkg/generics"
)

//Generator returns a function for building generators of the specified type
func Generator(mapType string) generics.GeneratorFactory {
	switch mapType {
	case "flat":
		return flat.NewGenerator
	}

	return flat.NewGenerator
}
