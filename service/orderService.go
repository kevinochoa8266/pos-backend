package service

import (
	"database/sql"

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

func (orderService *OrderService) CreateOrder(order []models.Order) bool {
	return false
}
