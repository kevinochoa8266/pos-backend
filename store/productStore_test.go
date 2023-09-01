package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var productStore = store.NewProductStore(db)

var product = models.Product{
	Id:        0,
	Name:      "Chocolate",
	Price:     5.00,
	Inventory: 100,
	StoreId:   2,
}

func TestSaveProduct(t *testing.T) {

	_, err := productStore.Save(&product)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetProduct(t *testing.T) {
	productStore.Save(&product)

	product, err := productStore.Get(1); if err != nil {
		t.Errorf("could not get product to to err: %s", err.Error())
	}
	
	if product.Name != "Chocolate" {
		t.Error("grabbed an unexpected product")
	}

	_, err = productStore.Get(10000)
	if err == nil {
		t.Error("product should not exist at id 1000")
	}
}
