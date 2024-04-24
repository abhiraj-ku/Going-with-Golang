package router

import (
	"example/go-post/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router:= new Router()

	// routes for all the CRUD
	router.HandleFunc("/api/stock/{id}",middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stock",middleware.GetAllStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/newStock",middleware.CreateStock).Methods("POST","OPTIONS")
	router.HandleFunc("/api/stock/{id}",middleware.UpdateStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deleteStock/{id}",middleware.DeleteStock).Methods("DELETE","OPTIONS")
}