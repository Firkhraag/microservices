package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/firkhraag/product-images/internal/handlers"
	"github.com/firkhraag/product-images/internal/util"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	// Bind address: :9090
	bindAddr = util.Env("BIND_ADDR", ":9090")
	// Log level: [debug, info, trace]
	logLevel = util.Env("LOG_LEVEL", "debug")
	// Base path to saved images: tmp/images
	basePath = util.Env("BASE_PATH", "tmp/images")
)

func main() {
	// Logger
	l := log.New(os.Stdout, "product-images", log.LstdFlags)

	// Local storage
	// max filesize 5MB
	store, err := files.NewLocal(*basePath, 1024*1000*5)
	if err != nil {
		l.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

	// Dependency injection of logger in files handler
	fh := handlers.NewFiles(l, store)
	// New router
	sm := mux.NewRouter()

	// CORS
	sm.Use(ghandlers.CORS(ghandlers.AllowedOrigins([]string{"http://localhost:8080"})))

	// GET endpoints
	getRouter := r.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	// POST endpoints
	postRouter := r.Methods("POST").Subrouter()
	postRouter.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.ServeHTTP)

	s := &http.Server{
		Addr:         bindAddr,
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
