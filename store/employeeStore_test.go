package store_test

import (
	"fmt"
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var db, _ = store.GetConnection(dbUrl)

var name = "John"
var number = 9417160432
var address = "123 abc"

func init() {
	err := store.CreateSchema(db)
	if err != nil {
		panic(err)
	}
	storeQuery := "INSERT INTO store (name, address) VALUES(?, ?)"

	_, err = db.Exec(storeQuery, "hello", "world")
	if err != nil {
		panic(err)
	}

	employeeQuery := "INSERT INTO employee (fullName, phoneNumber, address, storeId) VALUES(?, ?, ?, ?);"

	for i := 0; i < 3; i++ {
		_, err := db.Exec(employeeQuery, name, number, address, 1)
		if err != nil {
			panic(err)
		}
	}
}

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
	fmt.Println(employees)
}

func TestSave(t *testing.T) {
	employeeStore := store.NewEmployeeStore(db)

	employee := models.Employee{
		Id:      0,
		Name:    name,
		Phone:   number,
		Address: address,
		StoreId: 1,
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

	employee.Name = "Anthony"

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
