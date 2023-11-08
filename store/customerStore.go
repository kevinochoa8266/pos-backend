package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kevinochoa8266/pos-backend/models"
)

type CustomerStore struct {
	db *sql.DB
}

func NewCustomerStore(db *sql.DB) *CustomerStore {
	return &CustomerStore{db: db}
}

func (customerStore *CustomerStore) Save(customer *models.Customer) (string, error) {
	query := `INSERT INTO customer (
			id,
			firstName,
			lastName,
			phoneNumber,
			email,
			address
			)
			VALUES(?, ?, ?, ?, ?, ?);
	`

	result, err := customerStore.db.Exec(query, &customer.Id, &customer.FirstName, &customer.LastName, &customer.PhoneNumber, &customer.Email, &customer.Address)
	if err != nil {
		return "", fmt.Errorf("query failed to save customer to database, errors: %s", err.Error())
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		return "", errors.New("failed to insert reader into database")
	}
	return customer.Id, nil
}