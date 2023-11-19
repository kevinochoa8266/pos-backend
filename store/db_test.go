package store_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var dbUrl = ":memory:"
var name = "John"
var number = "9417160432"
var address = "123 abc"

var storeId = "FF"
var storeAddress = "123 abc street"
var storeCity = "Medellin"
var storeState = "Antioquia"
var storeCountry = "CO"
var storePostal = "050037"
var storeName = "XYZ Store"

var productName = "Chocolate"
var price int64 = 5
var inventory = 100

var db, _ = store.GetConnection(dbUrl)

func init() {
	if _, inCI := os.LookupEnv("GITHUB_ACTIONS"); inCI {
		fmt.Println("WE ARE IN THE ACTIONS ENV")
	} else {
		fmt.Println("WE ARE NOT IN ACTIONS ENV")
	}
	err := store.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	query := `INSERT INTO store (
				Id,
				address,
				city,
				state,
				country,
				postal,
				name
				)
				VALUES(?, ?, ?, ?, ?, ?, ?);
	`
	_, err = db.Exec(query, storeId, storeAddress, storeCity, storeState, storeCountry, storePostal, storeName)
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

	customerStore := store.NewCustomerStore(db)
	_, err = customerStore.Save(
		&models.Customer{
			Id:          "cu-123",
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "305-687-4999",
			Email:       "john.doe@gmail.com",
			Address:     "123 AVE",
		})

	if err != nil {
		panic(err)
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
