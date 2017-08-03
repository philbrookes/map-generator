package generics


//GeneratorFactory defines a factory function for generators
type GeneratorFactory func(topLeft, bottomRight Position, positionFactory PositionFactory, tileFactory TileFactory, mapFactory MapFactory) Generator

//PositionFactory defines a factory function for positions
type PositionFactory func(x, y, z Coordinate) Position

//TileFactory defines a factory function for positions
type TileFactory func(position Position) Tile

//MapFactory defines a factory function for maps
type MapFactory func(topLeft, bottomRight Position) Map

//MetadataFactory defines a factory function for metadata
type MetadataFactory func(float32, bool) Metadata