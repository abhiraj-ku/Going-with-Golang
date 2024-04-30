package router

import (
	"example/go-post/middlewares"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// routes for all the CRUD
	router.HandleFunc("/api/stock/{id}", middlewares.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middlewares.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newStock", middlewares.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middlewares.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteStock/{id}", middlewares.DeleteStock).Methods("DELETE", "OPTIONS")

	return router
}
