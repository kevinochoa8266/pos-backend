package service

import (
	"fmt"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
)

func ProccessCustomer(c models.Customer, customerStore *store.CustomerStore) (string, error) {
	var fullName = c.FirstName + " " + c.LastName

	params := &stripe.CustomerParams{
		Name:  &fullName,
		Phone: &c.PhoneNumber,
		Email: &c.Email,
		Address: &stripe.AddressParams{
			Line1: stripe.String(c.Address),
		},
	}

	client, err := customer.New(params)

	if err != nil {
		return "", fmt.Errorf("unable to create a new customer with stripe parameters: %v", err.Error())
	}

	// consult with anthony about middle names.
	storedCustomer := models.Customer{
		Id:          client.ID,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		PhoneNumber: client.Phone,
		Email:       client.Email,
		Address:     client.Address.Line1,
	}

	customerId, err := customerStore.Save(&storedCustomer)

	if err != nil {
		return "", err
	}

	return customerId, nil
}
