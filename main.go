package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"

	handlers "github.com/DragonSSS/cloud-audition-interview/handlers"
)

func main() {

	router := mux.NewRouter()

	// add handlers
	router.HandleFunc("/messages", handlers.CreateMessage).Methods("POST")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
