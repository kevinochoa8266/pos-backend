package store_test

import (
	"strconv"
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var dbUrl = ":memory:"
var name = "John"
var number = "9417160432"
var address = "123 abc"

var storeName = "XYZ Store"
var storeAddress = "123 abc street"

var productName = "Chocolate"
var price int64 = 5
var inventory = 100
var storeId = "FF"

var db, _ = store.GetConnection(dbUrl)

func init() {
	err := store.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	query := `INSERT INTO store (
				Id,
				name,
				address
				)
				VALUES(?, ?, ?);
	`
	_, err = db.Exec(query, storeId, storeName, storeAddress)
	if err != nil {
		panic(err)
	}

	employeeQuery := "INSERT INTO employee (fullName, phoneNumber, address, storeId) VALUES(?, ?, ?, ?);"

	for i := 0; i < 3; i++ {
		_, err := db.Exec(employeeQuery, name, number, address, storeId)
		if err != nil {
			panic(err)
		}
	}
	ps := store.NewProductStore(db)

	for i := 1; i < 10; i++ {
		_, err := ps.Save(&models.Product{
			Id:            strconv.Itoa(i),
			Name:          productName,
			UnitPrice:     price,
			Inventory:     inventory,
			BulkPrice:     price * 5,
			ItemsInPacket: 10,
			StoreId:       storeId})
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
