package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kevinochoa8266/pos-backend/handlers"
	"github.com/kevinochoa8266/pos-backend/utils"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not load env variables. err: %s", err.Error())
	}

	db_URL := os.Getenv("DB_URL")
	if err := utils.ReadCsvData("candy_data.csv", db_URL); err != nil {
		log.Println("File already exists")
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
