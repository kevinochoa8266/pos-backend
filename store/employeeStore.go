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

func (employeeStore *EmployeeStore) Save(employee *models.Employee) (int64, error) {
	query := `INSERT INTO employee (
				fullName,
				phoneNumber,
				address,
				storeId
				)
				VALUES(?, ?, ?, ?);
	`
	result, err := employeeStore.db.Exec(query, employee.Name, employee.Phone, employee.Address, employee.StoreId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (employeeStore *EmployeeStore) Get(id int) (*models.Employee, error) {
	query := `
		SELECT * FROM employee e WHERE e.id = ?;
	`
	row := employeeStore.db.QueryRow(query, id)
	if row.Err() != nil {
		return nil, errEmployeeNotFound
	}
	employee := models.Employee{}

	err := row.Scan(&employee.Id, &employee.Name, &employee.Phone, &employee.Address, &employee.StoreId)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (employeeStore *EmployeeStore) GetAll() ([]models.Employee, error) {
	query := `
		SELECT * FROM employee;
	`
	rows, err := employeeStore.db.Query(query)
	if err != nil {
		return nil, err
	}
	employees := []models.Employee{}

	for rows.Next() {
		employee := models.Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Phone, &employee.Address, &employee.StoreId)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (employeeStore *EmployeeStore) Update(employee *models.Employee) error {
	query := `
		UPDATE employee
		SET fullname = ?, phoneNumber = ?, address = ?, storeId = ?
		WHERE id = ? 
	`
	result, err := employeeStore.db.Exec(query, &employee.Name, &employee.Phone,
		employee.Address, employee.StoreId, employee.Id)
	if err != nil {
		return err
	}
	rowsUpdated, err := result.RowsAffected()
	if err != nil || rowsUpdated != 1 {
		return err
	}
	return nil
}

func (employeeStore *EmployeeStore) Delete(id int) error {
	query := `
		DELETE FROM employee WHERE id = ?;
	`
	result, err := employeeStore.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil || rowsDeleted != 1 {
		return errors.New("employee with the given id does not exist")
	}
	return nil
}
