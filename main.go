package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	handlers "github.com/DragonSSS/cloud-audition-interview/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

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

	service := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Info("Server started")
	<-done
	log.Info("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := service.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Info("Server Exited Properly")
}
