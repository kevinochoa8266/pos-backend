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
	query := "INSERT INTO orders (id, date, quantity, priceAtPurchase, productId, customerId) VALUES(?,?,?,?,?,?)"

	result, err := os.db.Exec(query,
		order.Id,
		order.Date,
		order.Quantity,
		order.PriceAtPurchase,
		order.ProductId,
		checkCustomerId(int16(order.CustomerId)),
	)
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
	query := "SELECT id, date, quantity, priceAtPurchase, productId, customerId FROM orders"

	rows, err := os.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("unable to perform query to get all orders. err: %s", err.Error())
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}

		var customerId sql.NullInt16

		err := rows.Scan(&order.Id, &order.Date,
			&order.Quantity, &order.PriceAtPurchase, &order.ProductId, &customerId)
		if err != nil {
			return nil, fmt.Errorf("unable to parse a row. err: %s", err.Error())
		}

		if customerId.Valid {
			order.CustomerId = int(customerId.Int16)
		} else {
			order.CustomerId = 0
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func (os *OrderStore) GetOrder(id string) ([]models.Order, error) {
	query := "SELECT id, date, quantity, priceAtPurchase, productId, customerId FROM orders WHERE id = ?"
	rows, err := os.db.Query(query, id)

	if err != nil {
		return nil, fmt.Errorf("unable to perform query to get all orders. err: %s", err.Error())
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}
		var customerId sql.NullInt16

		err := rows.Scan(&order.Id, &order.Date,
			&order.Quantity, &order.PriceAtPurchase, &order.ProductId, &customerId)
		if err != nil {
			return nil, fmt.Errorf("unable to parse a row. err: %s", err.Error())
		}

		if customerId.Valid {
			order.CustomerId = int(customerId.Int16)
		} else {
			order.CustomerId = 0
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func checkCustomerId(customerId int16) sql.NullInt16 {
	if customerId == 0 {
		return sql.NullInt16{}
	}
	return sql.NullInt16{
		Int16: customerId,
		Valid: true,
	}
}
