package store_test

import (
	"fmt"
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

func TestCustomerSave(t *testing.T) {
	newStore := store.NewCustomerStore(db)

	customer := models.Customer{
		Id:         "cu-124",
		FirstName: "alex",
		LastName: "Doe",
		PhoneNumber: "305-687-4999",
		Email: "alex.doe@gmail.com",
		Address: "123 AVE",
	}

	_, err := newStore.Save(&customer)
	if err != nil {
		t.Error("could not save a new customer in the database for testing")
	}
}

func TestCustomerGetByEmail(t *testing.T) {
	newStore := store.NewCustomerStore(db)

	customerEmail := "john.doe@gmail.com"

	id, err := newStore.GetByEmail(customerEmail)

	if err != nil {
		t.Errorf("could not retrieve a customer with email: %s", customerEmail)
	}

	knownId := "cu-123"

	if id != knownId {
		t.Errorf("incorrect customer retrieved, expected: %s, actual: %s", knownId, id)
	}

	customerEmail = "bob.doe@gmail.com"

	_, err = newStore.GetByEmail(customerEmail)
	if err == nil {
		t.Errorf("was not supposed to retrieve a customer with the email: %s", customerEmail)
	}

}

func TestCustomerGet(t *testing.T) {
	newStore := store.NewCustomerStore(db)

	customerId := "cu-123"

	_, err := newStore.Get(customerId)

	if err != nil {
		t.Errorf("could not retrieve a customer with id: %s", customerId)
	}

	customerId = "cu-256"

	_, err = newStore.Get(customerId)
	if err == nil {
		t.Errorf("was not supposed to retrieve a customer with the id: %s", customerId)
	}

	fmt.Println("HELLO THIS IS A TEST")

}