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
