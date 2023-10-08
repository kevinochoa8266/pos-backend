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
	image, err := getImage(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(err.Error())
	}
	if err := imageStore.Save(image); err != nil {
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
	images, err := imageStore.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(images)
}

func HandleUpdateImage(w http.ResponseWriter, r *http.Request) {
	image, err := getImage(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(err.Error())
		return
	}

	if err := imageStore.Update(image); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func HandleDeleteImage(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	deletedId, err := imageStore.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	message := make(map[string]string)
	message["id"] = deletedId
	json.NewEncoder(w).Encode(message)
}

func getImage(r *http.Request) (*models.Image, error) {
	id := mux.Vars(r)["id"]
	err := r.ParseMultipartForm(20 << 30) // This means 20 MB limit
	if err != nil {
		return nil, err
	}
	file, _, err := r.FormFile("image")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	imageContent, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return &models.Image{Id: id, Data: imageContent}, nil
}
