package Factory

import (
	"github.com/philbrookes/map-generator/pkg/Flat"
	"github.com/philbrookes/map-generator/pkg/Generics"
)

//Generator returns a function for building generators of the specified type
func Generator(mapType string) Generics.GeneratorFactory {
	switch mapType {
	case "flat":
		return Flat.NewGenerator
	}

	return Flat.NewGenerator
}
