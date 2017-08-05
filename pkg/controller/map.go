package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/philbrookes/map-generator/pkg/config"
	"github.com/philbrookes/map-generator/pkg/factory"
	"github.com/philbrookes/map-generator/pkg/generics"
)

//ConfigureMap configures the passed in router for handling map routes
func ConfigureMap(r *mux.Router, c *config.Config) {
	r.HandleFunc("/generate/{type}", generateMapHandlerFactory(c))
	r.HandleFunc("/generate/{type}/{x}/{y}", generateMapCoordHandlerFactory(c))
}

func generateMapHandlerFactory(c *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		mapType := vars["type"]

		w.Header().Set("Content-Type", "application/json")

		genMap, err := doGenerate(0, 0, mapType, c)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		json.NewEncoder(w).Encode(genMap)
	}

}
func generateMapCoordHandlerFactory(c *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		mapType := vars["type"]
		x, err := strconv.Atoi(vars["x"])
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		y, err := strconv.Atoi(vars["y"])
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		genMap, err := doGenerate(x, y, mapType, c)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(genMap)
	}

}

func doGenerate(x, y int, mapType string, c *config.Config) (generics.Map, error) {

	generatorFactory := factory.Generator(mapType)
	positionFactory := factory.Position(mapType)
	tileFactory := factory.Tile(mapType)
	mapFactory := factory.Map(mapType)

	buffer := c.GetBufferSize()
	viewport := c.GetViewportSize()
	mapsize := buffer + viewport

	topLeft := positionFactory(generics.Coordinate(x-(mapsize)), generics.Coordinate(y-(mapsize)), 0)
	bottomRight := positionFactory(generics.Coordinate(x+(mapsize)), generics.Coordinate(y+(mapsize)), 0)

	generator := generatorFactory(topLeft, bottomRight, positionFactory, tileFactory, mapFactory)

	return generator.Generate()

}
