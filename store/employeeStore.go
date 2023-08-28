package store

import (
	"database/sql"
	"errors"

	"github.com/kevinochoa8266/pos-backend/models"
)

var errEmployeeNotFound = errors.New("employee with id is not found")

type EmployeeStore struct {
	db *sql.DB
}

func NewEmployeeStore(db *sql.DB) *EmployeeStore {
	return &EmployeeStore{db: db}
}

func (employeeStore *EmployeeStore) Get(id int) (*models.Employee, error) {
	query := `
		SELECT * FROM employee e WHERE e.id = ?;
	`
	row := employeeStore.db.QueryRow(query, id); if row.Err() != nil {
		return nil, errEmployeeNotFound
	}
	 employee := models.Employee{}

	 err := row.Scan(&employee.Id, &employee.Name, &employee.Phone, &employee.Phone); if err != nil {
		return nil, err
	 }
	 return &employee, nil
}
