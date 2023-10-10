package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
	"github.com/stripe/stripe-go/v75/terminal/reader"
	readertesthelpers "github.com/stripe/stripe-go/v75/testhelpers/terminal/reader"
)

func HandleRegisterReader(w http.ResponseWriter, r *http.Request) {
	var req struct {
		LocationID string `json:"location_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	params := &stripe.TerminalReaderParams{
		Location:         stripe.String(req.LocationID),
		RegistrationCode: stripe.String("simulated-wpe"),
	}

	reader, err := reader.New(params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("reader.New: %v", err)
		return
	}

	WriteJSON(w, reader)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		PaymentIntentAmount string `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	amount, _ := strconv.ParseInt(req.PaymentIntentAmount, 10, 64)

	// For Terminal payments, the 'payment_method_types' parameter must include
	// 'card_present'.
	// To automatically capture funds when a charge is authorized,
	// set `capture_method` to `automatic`.
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card_present",
		}),
		CaptureMethod: stripe.String("manual"),
	}
	pi, err := paymentintent.New(params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("pi.New: %v", err)
		return
	}

	WriteJSON(w, pi)
}

func HandleProcessPayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ReaderID        string `json:"reader_id"`
		PaymentIntentID string `json:"payment_intent_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	params := &stripe.TerminalReaderProcessPaymentIntentParams{
		PaymentIntent: stripe.String(req.PaymentIntentID),
	}

	reader, err := reader.ProcessPaymentIntent(req.ReaderID, params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("reader.New: %v", err)
		return
	}

	WriteJSON(w, reader)
}

func HandleSimulatePayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ReaderID string `json:"reader_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	params := &stripe.TestHelpersTerminalReaderPresentPaymentMethodParams{}
	reader, err := readertesthelpers.PresentPaymentMethod(req.ReaderID, params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("reader.New: %v", err)
		return
	}

	WriteJSON(w, reader)
}

func HandleCapture(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		PaymentIntentID string `json:"payment_intent_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	pi, err := paymentintent.Capture(req.PaymentIntentID, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("pi.Capture: %v", err)
		return
	}

	WriteJSON(w, pi)
}

func WriteJSON(w http.ResponseWriter, v interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("io.Copy: %v", err)
		return
	}
}
