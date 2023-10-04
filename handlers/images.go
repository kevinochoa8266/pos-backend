package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var imageStore = store.NewImageStore(db)

func HandleSaveImage(w http.ResponseWriter, r *http.Request) {
	image := models.Image{}
	json.NewDecoder(r.Body).Decode(&image)

	if err := imageStore.Save(&image); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message := make(map[string]string)
		message["error"] = fmt.Sprintf("unable to save image with id: %s into the db", image.Id)
		json.NewEncoder(w).Encode(message)
	}
	w.WriteHeader(http.StatusOK)
}

func HandleGetImages(w http.ResponseWriter, r *http.Request) {
}
