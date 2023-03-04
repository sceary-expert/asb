package main

import (
	"log"
	"net/http"

	"example.com/configs" //add this
	"example.com/routes"  //add this

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router) //add this
	log.Fatal(http.ListenAndServe(":6000", router))
}
