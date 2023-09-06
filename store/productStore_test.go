package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var product = models.Product{
	Id:        "1024",
	Name:      "Chocolate",
	Price:     5.00,
	Inventory: 100,
	StoreId:   2,
}

var productStore = store.NewProductStore(db)

func TestSaveProduct(t *testing.T) {

	_, err := productStore.Save(&product)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetProduct(t *testing.T) {
	productStore.Save(&product)

	product, err := productStore.Get(product.Id)
	if err != nil {
		t.Errorf("could not get product to to err: %s", err.Error())
	}

	if product.Name != "Chocolate" {
		t.Error("grabbed an unexpected product")
	}

	_, err = productStore.Get("10000")
	if err == nil {
		t.Error("product should not exist at id 1000")
	}
}

func TestGetAllProducts(t *testing.T) {
	products, err := productStore.GetAll()
	if err != nil {
		t.Errorf("could not get all products: %s", err.Error())
	}

	if len(products) == 0 {
		t.Error("returned an empty slice when expecting products to return.")
	}
}

func TestUpdateProduct(t *testing.T) {

	candy, err := productStore.Get("1")
	if err != nil {
		t.Error("could not get the product from the database")
	}
	candy.Name = "Dulce de leche"

	err = productStore.Update(candy)
	if err != nil {
		t.Error("could not update the product.")
	}

	savedProduct, err := productStore.Get(candy.Id)
	if savedProduct.Name != candy.Name || err != nil {
		t.Errorf("could not update the product successfully")
	}

	// Change the id to where the product should not be able to update, since it does not exist.
	candy.Id = "5000"
	err = productStore.Update(candy)
	if err == nil {
		t.Error("expected an error updating a product with an id that does not exist")
	}
}

func TestDeleteProduct(t *testing.T) {
	validId := "2"
	invalidId := "2525"
	err := productStore.Delete(validId)
	if err != nil {
		t.Errorf("Could not delete product with id: %s, err: %s", validId, err.Error())
	}

	err = productStore.Delete(invalidId)
	if err == nil {
		t.Errorf("product with id %s should not exist", invalidId)
	}
}
