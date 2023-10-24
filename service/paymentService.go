package service

import (
	"fmt"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
	"github.com/stripe/stripe-go/v75/terminal/reader"
	readertesthelpers "github.com/stripe/stripe-go/v75/testhelpers/terminal/reader"
)

func TransactionProcess(params *stripe.PaymentIntentParams, payment models.Payment) (string, error) {

	pi, err := paymentintent.New(params)

	if err != nil {
		return "", fmt.Errorf("unable to create a new payment intent, error: %s", err.Error())
	}

	err = processPayment(payment.ReaderId, pi.ID)

	if err != nil {
		return "", err
	}

	resp, err := simulatePayment(payment.ReaderId)

	if err != nil {
		return "", err
	}
	return resp, nil
}

func processPayment(readerId string, paymentIntentId string) error {

	params := &stripe.TerminalReaderProcessPaymentIntentParams{
		PaymentIntent: stripe.String(paymentIntentId),
	}

	_, err := reader.ProcessPaymentIntent(readerId, params)

	if err != nil {
		return fmt.Errorf("the reader: %s was unable to process the payment inetent: %s, error: %s", readerId, paymentIntentId, err.Error())
	}
	return nil
}

func simulatePayment(readerId string) (string, error) {

	params := &stripe.TestHelpersTerminalReaderPresentPaymentMethodParams{
	}
	resp, err := readertesthelpers.PresentPaymentMethod(readerId, params)

	if err != nil {
		return "", fmt.Errorf("the reader: %s was unable to simulate the payment, error: %s", readerId, err.Error())
	}
	return string(resp.Action.Status), nil
}
