package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kevinochoa8266/pos-backend/handlers"
	"github.com/kevinochoa8266/pos-backend/utils"
)

func main() {
	if _, err := os.Open("store.db"); err != nil {
		if err := utils.ReadCsvData("candy_data.csv", "store.db"); err != nil {
			panic(err)
		}
	}

	router := mux.NewRouter()
	router.HandleFunc("/products", handlers.HandleGetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.HandleGetProduct).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", jsonContentTypeMiddleWare(router)))
}

func jsonContentTypeMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
