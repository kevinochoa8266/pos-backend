package service

import (
	"fmt"
	"time"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/stripe/stripe-go/v76/terminal/reader"
	readertesthelpers "github.com/stripe/stripe-go/v76/testhelpers/terminal/reader"
)

/*
1. Fetch order functionality. Fetch by id and date? CustomerId?
2. handle when customerId is ""
*/

func TransactionProcess(payment models.Payment, order *store.OrderStore, productStore *store.ProductStore, customerStore *store.CustomerStore) (string, error) {
	params, id, err := CreatePaymentIntentParams(payment, customerStore)

	if err != nil {
		return "", err
	}

	pi, err := paymentintent.New(params)

	if err != nil {
		return "", fmt.Errorf("unable to create a new payment intent, error: %s", err.Error())
	}

	err = ProcessPayment(payment.ReaderId, pi.ID)

	if err != nil {
		return "", err
	}

	resp, err := simulatePayment(payment.ReaderId)

	if err != nil {
		return "", err
	}

	if resp == "succeeded" {
		if id != "" {
			err := SaveOrder(pi.ID, pi.Created, payment, order, id)
			if err != nil {
				return resp, err
			}
		}

		err := ProcessInventory(payment, productStore)
		if err != nil {
			return "", err
		}
	}

	return resp, nil
}

func ProcessPayment(readerId string, paymentIntentId string) error {

	params := &stripe.TerminalReaderProcessPaymentIntentParams{
		PaymentIntent: stripe.String(paymentIntentId),
	}

	_, err := reader.ProcessPaymentIntent(readerId, params)

	if err != nil {
		return fmt.Errorf("the reader: %s was unable to process the payment intent: %s, error: %s", readerId, paymentIntentId, err.Error())
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

func SaveOrder(paymentId string, date int64, payment models.Payment, orderStore *store.OrderStore, customerId string) error {
	orderLen := len(payment.Products)

	for i := 0; i < orderLen; i++ {
		newOrder := models.Order{
			Id:              paymentId,
			ProductId:       payment.Products[i].ProductId,
			CustomerId:      customerId,
			Date:            time.Unix(date, 0),
			Quantity:        payment.Products[i].Quantity,
			PriceAtPurchase: payment.Products[i].Price,
		}
		err := orderStore.Save(&newOrder)
		if err != nil {
			return fmt.Errorf("unable to save order with id: %s into the database, error: %s", paymentId, err.Error())
		}
	}
	return nil
}

func ProcessInventory(payment models.Payment, productStore *store.ProductStore) error {
	orderLen := len(payment.Products)

	for i := 0; i < orderLen; i++ {
		product, err := productStore.Get(payment.Products[i].ProductId)
		if err != nil {
			return fmt.Errorf("unable to get product with id: %s while processing inventory, error: %s", payment.Products[i].ProductId, err.Error())
		}
		if payment.Products[i].BoughtInBulk {
			orderQuantity := payment.Products[i].Quantity * product.ItemsInPacket
			product.Inventory -= orderQuantity
		} else {
			product.Inventory -= payment.Products[i].Quantity
		}
		if err := productStore.Update(product); err != nil {
			return fmt.Errorf("could not update a product with a new inventory, err: %s", err.Error())
		}
	}
	return nil
}

func CreatePaymentIntentParams(payment models.Payment, customerStore *store.CustomerStore) (*stripe.PaymentIntentParams, string, error) {
	params := &stripe.PaymentIntentParams{}
	var id string

	if payment.CustomerEmail == "" {
		params = &stripe.PaymentIntentParams{
			Amount:   &payment.OrderTotal,
			Currency: stripe.String(string(stripe.CurrencyUSD)),
			PaymentMethodTypes: stripe.StringSlice([]string{
				"card_present",
			}),
			CaptureMethod: stripe.String("automatic"),
		}
	} else {
		var err error
		id, err = customerStore.GetByEmail(payment.CustomerEmail)

		if err != nil {
			params = &stripe.PaymentIntentParams{
				Amount:       &payment.OrderTotal,
				Currency:     stripe.String(string(stripe.CurrencyUSD)),
				ReceiptEmail: stripe.String(payment.CustomerEmail),
				PaymentMethodTypes: stripe.StringSlice([]string{
					"card_present",
				}),
				CaptureMethod: stripe.String("automatic"),
			}
			return params, "", nil
		}

		params = &stripe.PaymentIntentParams{
			Amount:       &payment.OrderTotal,
			Currency:     stripe.String(string(stripe.CurrencyUSD)),
			Customer:     stripe.String(id),
			ReceiptEmail: stripe.String(payment.CustomerEmail),
			PaymentMethodTypes: stripe.StringSlice([]string{
				"card_present",
			}),
			CaptureMethod: stripe.String("automatic"),
		}
	}

	return params, id, nil
}
