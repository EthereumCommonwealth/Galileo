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

	http.ListenAndServe(":8080", r)
}