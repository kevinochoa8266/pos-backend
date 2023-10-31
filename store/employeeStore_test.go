package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

func TestGet(t *testing.T) {
	employeeStore := store.NewEmployeeStore(db)

	id1 := 1
	_, err := employeeStore.Get(id1)
	if err != nil {
		t.Errorf("could not retrieve a employee with id: %d", id1)
	}
	id2 := 150
	_, err = employeeStore.Get(id2)
	if err == nil {
		t.Errorf("was not supposed to retrieve a user with the id: %d", id2)
	}
}

func TestGetAll(t *testing.T) {
	employeeStore := store.NewEmployeeStore(db)

	employees, err := employeeStore.GetAll()
	if err != nil {
		t.Error("could not get employees from the database")
	}

	if len(employees) == 0 {
		t.Error("no employees were retrieved")
	}
}

func TestSaveEmployee(t *testing.T) {
	employeeStore := store.NewEmployeeStore(db)

	employee := models.Employee{
		Id:      0,
		FullName:    name,
		Phone:   number,
		Address: address,
		StoreId: "FF",
	}
	_, err := employeeStore.Save(&employee)
	if err != nil {
		t.Error("could not store given employee into the database")
	}
}

func TestUpdate(t *testing.T) {
	employeeStore := store.NewEmployeeStore(db)

	employee, err := employeeStore.Get(3)
	if err != nil {
		t.Error("could not get the employee from the db")
	}

	employee.FullName = "Anthony"

	err = employeeStore.Update(employee)
	if err != nil {
		t.Errorf("could not update employee with id: %d", employee.Id)
	}
}

func TestDelete(t *testing.T) {
	employeeStore := store.NewEmployeeStore(db)

	err := employeeStore.Delete(2)
	if err != nil {
		t.Error("expected employee with id 2 to be deleted from database")
	}

	err = employeeStore.Delete(100)
	if err == nil {
		t.Error("expected there to not exist an employee with the given id")
	}
}
