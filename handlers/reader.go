package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/terminal/reader"
)

func HandleRegisterReader(w http.ResponseWriter, r *http.Request) {
	shopStore := store.NewShopStore(db)
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
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	params := &stripe.TerminalReaderParams{
		Location:         stripe.String(storeId),
		RegistrationCode: stripe.String(req.RegistrationCode),
		Label:            stripe.String(req.Label),
	}

	reader, err := reader.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("reader.New: %v", err)
		return
	}
	readerErr := service.SaveReader(reader.ID, reader.Location.ID, reader.Label, db)
	if readerErr != nil {
		panic(err)
	}
}

func HandleGetReaders(writer http.ResponseWriter, request *http.Request) {
	readerStore := store.NewReaderStore(db)
	readers, err := readerStore.GetAll()
	if err != nil {
		logger.Error("could not retrieve readers from the database. %s", err.Error(), 1)
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(readers)
}
