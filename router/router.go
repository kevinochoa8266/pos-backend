package router

import (
	"github.com/gorilla/mux"
	"github.com/kevinochoa8266/pos-backend/handlers"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products/favorites", handlers.HandleGetFavorites).Methods("GET")
	router.HandleFunc("/products", handlers.HandleGetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.HandleGetProduct).Methods("GET")
	router.HandleFunc("/products", handlers.HandleAddProduct).Methods("POST")
	router.HandleFunc("/products", handlers.HandleUpdateProduct).Methods("PUT")
	router.HandleFunc("/products", handlers.HandleDeleteProduct).Methods("DELETE")

	router.HandleFunc("/images", handlers.HandleGetImages).Methods("GET")
	router.HandleFunc("/images/{id}", handlers.HandleSaveImage).Methods("POST")
	router.HandleFunc("/images/{id}", handlers.HandleUpdateImage).Methods("PUT")
	router.HandleFunc("/images/{id}", handlers.HandleDeleteImage).Methods("DELETE")
	router.HandleFunc("/readers", handlers.HandleRegisterReader).Methods("POST")
	router.HandleFunc("/readers", handlers.HandleGetReaders).Methods("GET")
	router.HandleFunc("/payments", handlers.HandleTransaction).Methods("POST")
	router.HandleFunc("/customers", handlers.HandleCreateCustomer).Methods("POST")

	return router
}
