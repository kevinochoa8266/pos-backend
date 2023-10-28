package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/stripe/stripe-go/v75"
)

var orderStore *store.OrderStore

func HandleTransaction(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error("unable to decode json body to create a payment intent, error: %s", err.Error(), 1)
		return
	}

	amount, _ := strconv.ParseInt(payment.OrderTotal, 10, 64)

	params := &stripe.PaymentIntentParams{
		Amount:       stripe.Int64(amount),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Customer:     stripe.String(payment.CustomerId),
		ReceiptEmail: stripe.String("kevin.ochoa@ufl.edu"), // this will be changed to the customer email
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card_present",
		}),
		CaptureMethod: stripe.String("automatic"),
	}
	response, err := service.TransactionProcess(params, payment, orderStore)

	if err != nil {
		logger.Error(err.Error())
	}

	if response == "succeeded" {
		json.NewEncoder(w).Encode("Transaction successful.")
	} else {
		json.NewEncoder(w).Encode("Transaction unsuccessful.")
	}
}
