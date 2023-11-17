package service_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/stripe/stripe-go/v75/customer"
)

func TestProcessCustomer(t *testing.T) {
	cust := models.Customer{
		FirstName:   "Jane",
		LastName:    "Doe",
		PhoneNumber: "305-888-8888",
		Email:       "jane.doe@test.com",
		Address:     "123 Main Street, City, Country",
	}

	fullName := cust.FirstName + " " + cust.LastName

	id, err := service.ProccessCustomer(cust, customerStore)

	if err != nil {
		t.Errorf("unable to create customer, error: %s", err)
	}

	c, _ := customer.Get(id, nil)

	if id != c.ID {
		t.Errorf("customer was not created correctly, expected: %s, actual: %s", id, c.ID)
	}

	if fullName != c.Name {
		t.Errorf("FirstName mismatch, expected: %s, actual: %s", fullName, c.Name)
	}

	if cust.Email != c.Email {
		t.Errorf("Email mismatch, expected: %s, actual: %s", cust.Email, c.Email)
	}

}
