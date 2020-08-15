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
	// create
	router.HandleFunc("/messages", handlers.CreateMessage).Methods("POST")
	// get
	router.HandleFunc("/messages/{messageID}", handlers.GetMessage).Methods("GET")
	// get-all
	router.HandleFunc("/messages", handlers.GetMessages).Methods("GET")
	// update
	router.HandleFunc("/messages/{messageID}", handlers.UpdateMessage).Methods("PUT")
	// delte
	router.HandleFunc("/messages/{messageID}", handlers.DeleteMessage).Methods("DELETE")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
