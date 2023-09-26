package store_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var os = store.NewOrderStore(db)

func TestSaveOrder(t *testing.T) {
	order := models.Order{
		Id: uuid.New(),
		ProductId:  "1",
		CustomerId: 0,
		Date:       time.Now(),
		Quantity:   5,
		TotalPrice: 4500,
	}

	err := os.Save(&order)
	if err != nil {
		t.Error("unable to save a order into the database")
	}

	order.ProductId = "id does not exist"
	err = os.Save(&order)
	if err == nil {
		t.Error("the productId associated with the order does not exist")
	}

}
