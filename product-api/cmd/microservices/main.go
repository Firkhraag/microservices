package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/firkhraag/product-api/internal/handlers"
	ghandlers "github.com/gorilla/handlers"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {

	// Bind addres of the service from the env variable
	bindAddress := os.Getenv("BIND_ADDRESS")
	// Default is :8080
	if bindAddress == "" {
		bindAddress = ":8080"
	}

	// Logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// Dependency injection of logger in products chandler
	ph := handlers.NewProducts(l)
	// New router
	sm := mux.NewRouter()

	// CORS
	sm.Use(ghandlers.CORS(ghandlers.AllowedOrigins([]string{"http://localhost:3000"})))

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProduct)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/products", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	opts := middleware.RedocOpts{
		SpecURL: "/swagger.yaml",
	}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	s := &http.Server{
		Addr:         bindAddress,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
