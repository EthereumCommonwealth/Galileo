package main

import (
	"github.com/gorilla/mux"
	"github.com/eduartua/callisto/Galileo/controllers"
	"net/http"
	"fmt"
	"github.com/eduartua/callisto/Galileo/models"
)

//This is going to be placed in a config file. For dev purposes will be here, then moved.
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hola"
	dbname   = "galileo_dev"
)

func main() {
	//Create db conn string. Used then to create the model services.
	//This is going to be placed in a config file. For dev purposes will be here, then moved.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	//Assets Handlers
	assetHandler := http.FileServer(http.Dir("./assets"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	http.ListenAndServe(":8080", r)
}