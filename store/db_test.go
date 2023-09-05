package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/store"
)

var dbUrl = ":memory:"
var name = "John"
var number = "9417160432"
var address = "123 abc"

var storeName = "XYZ Store"
var storeAddress = "123 abc street"

var db, _ = store.GetConnection(dbUrl)

func init() {
	err := store.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	storeQuery := "INSERT INTO store (name, address) VALUES(?, ?)"

	for i := 0; i < 3; i++ {
		_, err := db.Exec(storeQuery, storeName, storeAddress)
		if err != nil {
			panic(err)
		}
	}

	employeeQuery := "INSERT INTO employee (fullName, phoneNumber, address, storeId) VALUES(?, ?, ?, ?);"

	for i := 0; i < 3; i++ {
		_, err := db.Exec(employeeQuery, name, number, address, 1)
		if err != nil {
			panic(err)
		}
	}

}

func TestGetConnection(t *testing.T) {
	_, err := store.GetConnection(dbUrl)
	if err != nil {
		t.Error("could not establish connection to the database.")
	}
}

func TestCloseConnection(t *testing.T) {
	db, _ := store.GetConnection(dbUrl)
	err := store.CloseConnection(db)
	if err != nil {
		t.Error("could not close the database connection")
	}
}

func TestCreateSchema(t *testing.T) {
	db, _ := store.GetConnection(dbUrl)
	err := store.CreateSchema(db)
	if err != nil {
		t.Error("could not create the given schema.")
	}
}
