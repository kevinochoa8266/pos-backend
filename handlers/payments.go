package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/kevinochoa8266/pos-backend/store"
)

var orderStore *store.OrderStore

func HandleTransaction(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error("unable to decode json body to create a payment intent, error: %s", err.Error(), 1)
		return
	}

	response, err := service.TransactionProcess(payment, orderStore, productStore, customerStore)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		message := make(map[string]string)
		message["error"] = err.Error()
		json.NewEncoder(w).Encode(message)
		return
	}

	if response == "succeeded" {
		json.NewEncoder(w).Encode("Transaction successful.")
	} else {
		json.NewEncoder(w).Encode("Transaction unsuccessful.")
	}
}
