package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var database, _ = store.GetConnection(dbUrl)

var storeName = "XYZ Store"
var storeAddress = "123 abc street"

func init() {
	err := store.CreateSchema(database)
	if err != nil {
		panic(err)
	}
	storeQuery := "INSERT INTO store (name, address) VALUES(?, ?)"

	for i := 0; i < 3; i++ {
		_, err := database.Exec(storeQuery, storeName, storeAddress)
		if err != nil {
			panic(err)
		}
	}
}

func TestStoreSave(t *testing.T) {
	newStore := store.NewStore(database)

	store := models.Store{
		Name:    "ABC Store",
		Address: "345 AVE",
	}

	_, err := newStore.Save(&store)
	if err != nil {
		t.Error("could not save a new store in the database")
	}
}
