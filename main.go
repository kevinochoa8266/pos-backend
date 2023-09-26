package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinochoa8266/pos-backend/app"
	"github.com/kevinochoa8266/pos-backend/handlers"
)

func main() {

	err := app.SetupApp()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/products", handlers.HandleGetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.HandleGetProduct).Methods("GET")
	router.HandleFunc("/products", handlers.HandleAddProduct).Methods("POST")

	log.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleWare(router)))
}

func jsonContentTypeMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
