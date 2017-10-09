package factory

import (
	"github.com/philbrookes/map-generator/pkg/flat"
	"github.com/philbrookes/map-generator/pkg/generics"
)

// Metadata returns a MetadataFactory for creating new metadata in the supplied type
func Metadata(mapType string) generics.MetadataFactory {
	switch mapType {
	case "flat":
		return flat.NewMetadata
	}

	return flat.NewMetadata
}

