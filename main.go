package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	handlers "github.com/DragonSSS/cloud-audition-interview/handlers"
	messagemap "github.com/DragonSSS/cloud-audition-interview/messagemap"

	log "github.com/sirupsen/logrus"
)

var healthy int32

// @title Cloud Audition API
// @version 1.0
// @description This is a service to manage messages and check palindrome
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email demo@gamil.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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

	// healthz
	router.HandleFunc("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	}))

	service := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		atomic.StoreInt32(&healthy, 1)
		if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Info("Server started")
	<-done
	log.Info("Server Stopping")
	atomic.StoreInt32(&healthy, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		messagemap.CleanMap()
		cancel()
	}()

	if err := service.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Info("Server Exited Properly")
}
