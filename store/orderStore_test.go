package store_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var orderStore = store.NewOrderStore(db)

func TestSaveOrder(t *testing.T) {
	order := models.Order{
		Id:         "tr_jubilee",
		ProductId:  "1",
		CustomerId: 0,
		Date:       time.Now(),
		Quantity:   5,
	}

	err := orderStore.Save(&order)
	if err != nil {
		t.Errorf("unable to save a order into the database, err: %s", err.Error())
	}

	order.ProductId = "id does not exist"
	err = orderStore.Save(&order)
	if err == nil {
		t.Error("the productId associated with the order does not exist")
	}
}

func TestGetOrders(t *testing.T) {
	order := models.Order{
		Id:         "tr_jubo",
		ProductId:  "1",
		CustomerId: 0,
		Date:       time.Now(),
		Quantity:   5,
	}

	for i := 0; i < 2; i++ {
		order.Id = order.Id + strconv.Itoa(i)
		err := orderStore.Save(&order)
		if err != nil {
			t.Error("unable to create the orders for the test")
		}
	}

	orders, err := orderStore.GetOrders()
	if err != nil {
		t.Errorf("unable to get all orders. err: %s", err.Error())
	}
	if len(orders) == 0 {
		t.Errorf("no orders were returned. Total was %d", len(orders))
	}
}

func TestGetOrder(t *testing.T) {
	order := models.Order{
		Id:         "tr_abc123",
		Date: time.Now(),
		Quantity:   5,
		PriceAtPurchase: 10,
		ProductId: "1",
		CustomerId: 0,
	}

	for i := 0; i < 3; i++ {
		productId, _ := strconv.Atoi(order.ProductId)
		productId += 1
		order.ProductId = strconv.Itoa(productId)
		err := orderStore.Save(&order)
		if err != nil {
			t.Error("unable to save orders for test.")
		}
	}

	orders, err := orderStore.GetOrder(order.Id)
	if err != nil {
		t.Error(err.Error())
	}
	if len(orders) == 0 {
		t.Errorf("unable to grab the order with the given id: %s", order.Id)
	}

	orders, _ = orderStore.GetOrder("abc 123")
	if len(orders) != 0 {
		t.Error("invalid id should have returned 0 orders associated with it")
	}
}
