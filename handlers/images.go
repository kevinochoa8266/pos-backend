package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var imageStore *store.ImageStore

func HandleSaveImage(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := r.ParseMultipartForm(20 << 30) // This means 20 MB limit
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message := make(map[string]string)
		message["error"] = "The file was unable to be parsed and uploaded. Check if the file is greater than 20 MB"
		json.NewEncoder(w).Encode(message)
	}
	file, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(err.Error())
	}
	defer file.Close()

	imageContent, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error(err.Error())
	}

	image := models.Image{Id: id, Data: imageContent}
	if err := imageStore.Save(&image); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message := make(map[string]string)
		message["error"] = fmt.Sprintf("unable to save image with id: %s into the db", image.Id)
		json.NewEncoder(w).Encode(message)
		logger.Error(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func HandleGetImages(w http.ResponseWriter, r *http.Request) {
}
