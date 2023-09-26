package main

import (
	"log"
	"net/http"

	"github.com/kevinochoa8266/pos-backend/app"
	"github.com/kevinochoa8266/pos-backend/router"
)

func main() {
	// add changes to git
	err := app.SetupApp()
	if err != nil {
		panic(err)
	}

	router := router.CreateRouter()

	log.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleWare(router)))
}

func jsonContentTypeMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
