package store

import (
	"database/sql"
	"fmt"

	"github.com/kevinochoa8266/pos-backend/models"
)

type OrderStore struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) *OrderStore {
	return &OrderStore{db: db}
}

func (os *OrderStore) Save(order *models.Order) error {
	query := "INSERT INTO orders (id, productId, customerId, date, quantity, totalPrice) VALUES(?,?,?,?,?,?)"

	result, err := os.db.Exec(query,
		order.Id,
		order.ProductId,
		order.CustomerId,
		order.Date,
		order.Quantity,
		order.TotalPrice)
	if err != nil {
		return fmt.Errorf("unable to perform insert of order. err: %s", err.Error())
	}
	rows, err := result.RowsAffected()
	if rows != 1 || err != nil {
		return fmt.Errorf("unable to save the order to the database. err: %s", err.Error())
	}
	return nil
}

func (os *OrderStore) GetOrders() ([]models.Order, error) {
	query := "SELECT id, productId, customerId, date, quantity, totalPrice FROM orders"

	rows, err := os.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("unable to perform query to get all orders. err: %s", err.Error())
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}

		err := rows.Scan(&order.Id, &order.ProductId,
			&order.CustomerId, &order.Date, &order.Quantity, &order.TotalPrice)
		if err != nil {
			return nil, fmt.Errorf("unable to parse a row. err: %s", err.Error())
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (os *OrderStore) GetOrder(id string) ([]models.Order, error) {
	query := "SELECT id, productId, customerId, date, quantity, totalPrice FROM orders WHERE id = ?"
	rows, err := os.db.Query(query, id)

	if err != nil {
		return nil, fmt.Errorf("unable to perform query to get all orders. err: %s", err.Error())
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}

		err := rows.Scan(&order.Id, &order.ProductId,
			&order.CustomerId, &order.Date, &order.Quantity, &order.TotalPrice)
		if err != nil {
			return nil, fmt.Errorf("unable to parse a row. err: %s", err.Error())
		}
		orders = append(orders, order)
	}
	return orders, nil
}
