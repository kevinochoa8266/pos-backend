package service

import (
	"fmt"
	"time"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
	"github.com/stripe/stripe-go/v75/terminal/reader"
	readertesthelpers "github.com/stripe/stripe-go/v75/testhelpers/terminal/reader"
)

/*
1. Save the order in the DB once the payment is successful
   * The order table does not create boughtInBulk but the models.order does, ask tony.
2. Send email reciepts to customers
3. Update inventory after payment is successful
4. Create customer with stripe and save them in our db
5. Fetch order functionality. Fetch by id and date? CustomerId?
6. will there ever be refunds or cancellations?
7. check inventory while the order is being created.
8. handle when customerId is ""
*/

func TransactionProcess(payment models.Payment, order *store.OrderStore) (string, error) {
	params := &stripe.PaymentIntentParams{
		Amount:       &payment.OrderTotal,
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Customer:     stripe.String(payment.CustomerId),
		ReceiptEmail: stripe.String("kevin.ochoa@ufl.edu"), // this will be changed to the customer email
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card_present",
		}),
		CaptureMethod: stripe.String("automatic"),
	}

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

	if resp == "succeeded" {
		err = SaveOrder(pi.ID, pi.Created, payment, order)
		if err != nil {
			return resp, err //Let anthony know that I want the transaction to be successful even if the order is unable to save.
		}
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

	params := &stripe.TestHelpersTerminalReaderPresentPaymentMethodParams{}
	resp, err := readertesthelpers.PresentPaymentMethod(readerId, params)

	if err != nil {
		return "", fmt.Errorf("the reader: %s was unable to simulate the payment, error: %s", readerId, err.Error())
	}

	if resp.Action.Status == "failed" {
		return string(resp.Action.Status), fmt.Errorf("failure message: %s", resp.Action.FailureMessage)
	} else {
		return string(resp.Action.Status), nil
	}
}

func SaveOrder(paymentId string, date int64, payment models.Payment, orderStore *store.OrderStore) error {
	orderLen := len(payment.Products)

	for i := 0; i < orderLen; i++ {
		newOrder := models.Order{
			Id:                     paymentId,
			ProductId:              payment.Products[i].ProductId,
			CustomerId:             payment.CustomerId,
			Date:                   time.Unix(date, 0),
			Quantity:               payment.Products[i].Quantity,
			PriceAtPurchase: payment.Products[i].Price,
		}
		err := orderStore.Save(&newOrder)
		if err != nil {
			return err
		}
	}
	return nil
}
