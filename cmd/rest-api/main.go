package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/philbrookes/map-generator/pkg/config"
	"github.com/philbrookes/map-generator/pkg/controller"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	originsOk := handlers.AllowedOrigins(config.GetConfig().GetAllowedOrigins())
	methodsOk := handlers.AllowedMethods(config.GetConfig().GetAllowedMethods())

	controller.ConfigureMap(router.PathPrefix("/api/map/").Subrouter(), config.GetConfig())

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(config.GetConfig().GetPortListenerStr(), handlers.CORS(originsOk, methodsOk)(router)))
}
