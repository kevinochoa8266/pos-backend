package router

import (
	"github.com/gorilla/mux"
	"github.com/kevinochoa8266/pos-backend/handlers"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", handlers.HandleGetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.HandleGetProduct).Methods("GET")
	router.HandleFunc("/products", handlers.HandleAddProduct).Methods("POST")
	router.HandleFunc("/products", handlers.HandleUpdateProduct).Methods("PUT")
	router.HandleFunc("/products", handlers.HandleDeleteProduct).Methods("DELETE")

	return router
}
