package handlers

import (
	"net/http"

	"github.com/kevinochoa8266/pos-backend/store"
)

var imageStore = store.NewImageStore(db)

func HandleGetImages(w http.ResponseWriter, r *http.Request) {
	
}
