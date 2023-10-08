package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/kevinochoa8266/pos-backend/app"
	"github.com/kevinochoa8266/pos-backend/router"
)

func main() {
	err := app.SetupApp()
	if err != nil {
		panic(err)
	}

	router := router.CreateRouter()

	log.Println("Starting server at port 8080")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:4200"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"}),
	)(router)))
}
