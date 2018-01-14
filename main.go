package main

import (
	"github.com/gorilla/mux"
	"github.com/eduartua/callisto/Galileo/controllers"
	"net/http"
)

func main() {
	staticC := controllers.NewStatic()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")

	//Assets Handlers
	assetHandler := http.FileServer(http.Dir("./assets"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	http.ListenAndServe(":8080", r)
}