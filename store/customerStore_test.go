package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

func TestCustomerSave(t *testing.T) {
	newStore := store.NewCustomerStore(DB)

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