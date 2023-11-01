package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/stripe/stripe-go/v75"
)

var shopStore *store.ShopStore
var readerStore *store.ReaderStore

func HandleRegisterReader(w http.ResponseWriter, r *http.Request) {
	stores, err := shopStore.GetAll()
	if err != nil {
		panic(err)
	}
	storeId := stores[0].Id

	var req struct {
		RegistrationCode string `json:"registration_code"`
		Label            string `json:"label"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error("unable to decode json body to register a reader, error: %s", err.Error(), 1)
		return
	}

	params := &stripe.TerminalReaderParams{
		Location:         stripe.String(storeId),
		RegistrationCode: stripe.String(req.RegistrationCode),
		Label:            stripe.String(req.Label),
	}

	readerId, err := service.SaveReader(params, readerStore)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(readerId)
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
