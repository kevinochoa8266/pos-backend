package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/kevinochoa8266/pos-backend/store"
)

var shopStore *store.ShopStore
var readerStore *store.ReaderStore

func HandleRegisterReader(w http.ResponseWriter, r *http.Request) {
	var Req struct {
		RegistrationCode string `json:"registration_code"`
		Label            string `json:"label"`
	}

	if err := json.NewDecoder(r.Body).Decode(&Req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error("unable to decode json body to register a reader, error: %s", err.Error(), 1)
		return
	}
	err := service.SaveReader(Req, readerStore, shopStore)

	if err != nil {
		logger.Error(err.Error())
	}
}

func HandleGetReaders(writer http.ResponseWriter, request *http.Request) {
	readers, err := readerStore.GetAll()
	if err != nil {
		logger.Error("could not retrieve readers from the database. %s", err.Error(), 1)
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(readers)
}