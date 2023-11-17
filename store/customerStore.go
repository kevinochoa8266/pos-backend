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

func (customerStore *CustomerStore) Get(id string) (*models.Customer, error) {
	query := `
		SELECT * FROM customer c WHERE c.id = ?;
	`
	row := customerStore.db.QueryRow(query, id)
	if row.Err() != nil {
		return nil, fmt.Errorf("unable to fetch customer with id: %s, error: %s", id, row.Err().Error())
	}
	customer := models.Customer{}

	err := row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.PhoneNumber, &customer.Email, &customer.Address)
	if err != nil {
		return nil, fmt.Errorf("unable to parse a row, err: %s", err.Error())
	}
	return &customer, nil
}

func (customerStore *CustomerStore) GetByEmail(email string) (string, error) {
	query := `
		SELECT * FROM customer c WHERE c.email = ?;
	`
	row := customerStore.db.QueryRow(query, email)
	if row.Err() != nil {
		return "", fmt.Errorf("unable to fetch customer with email: %s, error: %s", email, row.Err().Error())
	}
	customer := models.Customer{}

	err := row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.PhoneNumber, &customer.Email, &customer.Address)
	if err != nil {
		return "", fmt.Errorf("unable to parse a row, err: %s", err.Error())
	}
	return customer.Id, nil
}
