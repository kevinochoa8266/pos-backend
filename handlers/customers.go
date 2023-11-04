package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/kevinochoa8266/pos-backend/store"
)

var customerStore *store.CustomerStore

func HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error("unable to decode json body to create a customer, error: %s", err.Error(), 1)
		return
	}

	customerId, err := service.ProccessCustomer(customer, customerStore)

	if err != nil {
		logger.Error(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customerId)

}