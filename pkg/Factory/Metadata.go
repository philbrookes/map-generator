package Factory

import (
	"github.com/philbrookes/map-generator/pkg/Flat"
	"github.com/philbrookes/map-generator/pkg/Generics"
)

// Metadata returns a MetadataFactory for creating new metadata in the supplied type
func Metadata(mapType string) Generics.MetadataFactory {
	switch mapType {
	case "flat":
		return Flat.NewMetadata
	}

	return Flat.NewMetadata
}

