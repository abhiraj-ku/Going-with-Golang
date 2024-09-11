package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("Hello, world!")
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Environment variable PORT is not set")
	}

	// router
	// for closing the server on ctrl+c gracefully
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	router := chi.NewRouter()

	// Cors error
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://example.com", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Route to handle /ready
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// go channel to start the server

	go func() {
		fmt.Printf("server started on port: %v", port)
		srvOn := srv.ListenAndServe()

		if srvOn != nil {
			log.Fatal(srvOn)
		}
		fmt.Println("Port:", port)

	}()

	<-stopChan
	log.Println("Stopping the server gracefully")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()

	log.Println("server stopped gracefully")

}