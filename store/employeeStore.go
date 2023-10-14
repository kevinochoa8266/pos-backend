package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kevinochoa8266/pos-backend/models"
)

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
	result, err := employeeStore.db.Exec(query, &employee.Name, &employee.Phone, &employee.Address, &employee.StoreId)
	if err != nil {
		return 0, fmt.Errorf("unable to save employee into database, error: %s", err.Error())
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
		return nil, fmt.Errorf("unable to fetch employee with id: %d, error: %s", id, row.Err().Error())
	}
	employee := models.Employee{}

	err := row.Scan(&employee.Id, &employee.Name, &employee.Phone, &employee.Address, &employee.StoreId)
	if err != nil {
		return nil, fmt.Errorf("unable to parse a row, err: %s", err.Error())
	}
	return &employee, nil
}

func (employeeStore *EmployeeStore) GetAll() ([]models.Employee, error) {
	query := `
		SELECT * FROM employee;
	`
	rows, err := employeeStore.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch all employees from database, error: %s", err.Error())
	}
	defer rows.Close()
	employees := []models.Employee{}

	for rows.Next() {
		employee := models.Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Phone, &employee.Address, &employee.StoreId)
		if err != nil {
			return nil, fmt.Errorf("unable to parse a row, err: %s", err.Error())
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
		return fmt.Errorf("unable to update employee in database with id: %d, error: %s", employee.Id, err.Error())
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
		return fmt.Errorf("unable to delete employee in database with id: %d, error: %s", id, err.Error())
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil || rowsDeleted != 1 {
		return errors.New("employee with the given id does not exist")
	}
	return nil
}
