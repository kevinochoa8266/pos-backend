package store_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var os = store.NewOrderStore(db)

func TestSaveOrder(t *testing.T) {
	order := models.Order{
		Id:           uuid.New(),
		ProductId:    "1",
		CustomerId:   0,
		Date:         time.Now(),
		BoughtInBulk: false,
		Quantity:     5,
		TotalPrice:   4500,
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

func TestGetOrders(t *testing.T) {
	order := models.Order{
		Id:         uuid.New(),
		ProductId:  "1",
		CustomerId: 0,
		Date:       time.Now(),
		Quantity:   5,
		TotalPrice: 4500,
	}

	for i := 0; i < 2; i++ {
		order.Id = uuid.New()
		err := os.Save(&order)
		if err != nil {
			t.Error("unable to create the orders for the test")
		}
	}

	orders, err := os.GetOrders()
	if err != nil {
		t.Errorf("unable to get all orders. err: %s", err.Error())
	}
	if len(orders) == 0 {
		t.Errorf("no orders were returned. Total was %d", len(orders))
	}
}

func TestGetOrder(t *testing.T) {

	order := models.Order{
		Id:           uuid.New(),
		ProductId:    "6",
		CustomerId:   0,
		Date:         time.Now(),
		BoughtInBulk: false,
		Quantity:     5,
		TotalPrice:   4500,
	}

	for i := 0; i < 3; i++ {
		productId, _ := strconv.Atoi(order.ProductId)
		productId += 1
		order.ProductId = strconv.Itoa(productId)
		err := os.Save(&order)
		if err != nil {
			t.Error("unable to save orders for test.")
		}
	}

	orders, err := os.GetOrder(order.Id.String())
	if err != nil {
		t.Error(err.Error())
	}
	if len(orders) == 0 {
		t.Errorf("unable to grab the order with the given id: %s", order.Id)
	}

	orders, _ = os.GetOrder("abc 123")
	if len(orders) != 0 {
		t.Error("invalid id should have returned 0 orders associated with it")
	}
}
