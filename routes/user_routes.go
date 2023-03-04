package routes

import (
	"example.com/controllers" //add this

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	//All routes related to users comes here
	router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET") //add this
}
