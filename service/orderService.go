package service

import (
	"database/sql"
	"log/slog"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

type OrderService struct {
	orderStore   *store.OrderStore
	productStore *store.ProductStore
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		orderStore:   store.NewOrderStore(db),
		productStore: store.NewProductStore(db),
	}
}

func (orderService *OrderService) GetOrders() ([]models.Order, error) {
	orders, err := orderService.orderStore.GetOrders()
	if err != nil {
		slog.Error("Unable to retrieve the orders from the database", "err", err.Error())
		return nil, err
	}
	return orders, nil
}

func (orderService *OrderService) SaveOrder(order []models.Order) error {
	for _, orderValue := range order {
		if err := orderService.orderStore.Save(&orderValue); err != nil {
			slog.Error("Unable to save order to the database.", "err", err.Error())
			return err
		}
	}
	slog.Info("Successfully saved an order to the database")
	return nil
}

func (orderService *OrderService) GetOrder(id string) ([]models.Order, error) {
	orders, err := orderService.orderStore.GetOrder(id)
	if err != nil {
		slog.Error("Unable to retrieve the order with the given id.", "id", id, "err", err.Error())
	}
	return orders, nil
}
