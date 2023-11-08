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

	var customerId *sql.NullString

	if order.CustomerId == "0" {
		customerId = &sql.NullString{}
	}

	customerId = &sql.NullString{
		String: order.CustomerId,
	}

	result, err := os.db.Exec(query,
		order.Id,
		order.Date,
		order.Quantity,
		order.PriceAtPurchase,
		order.ProductId,
		customerId,
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

	orders, err := os.retrieveOrders(rows)
	if err != nil {
		return nil, err
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

	orders, err := os.retrieveOrders(rows)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (os *OrderStore) retrieveOrders(rows *sql.Rows) ([]models.Order, error) {
	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}

		customerId := sql.NullString{}

		err := rows.Scan(&order.Id, &order.Date,
			&order.Quantity, &order.PriceAtPurchase, &order.ProductId, &customerId)
		if err != nil {
			return nil, fmt.Errorf("unable to parse the rows for retrieving orders, err: %s", err.Error())
		}

		if !customerId.Valid {
			order.CustomerId = "0"
		} else {
			order.CustomerId = customerId.String
		}
		orders = append(orders, order)
	}
	return orders, nil
}
