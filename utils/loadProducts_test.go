package utils_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/kevinochoa8266/pos-backend/utils"
)

func TestLoadProducts(t *testing.T) {
	db, _ := store.GetConnection(":memory:")

	store.CreateSchema(db)
	shopStore := store.NewStore(db)
	id, err := shopStore.Save(&models.Store{Id: 1, Name: "testStore", Address: "123 abc"})
	if err != nil {
		t.Errorf("unable to set up db for testing.")
	}

	if err = utils.LoadProductsIntoStore(id, db, "../candy_data.csv"); err != nil {
		t.Errorf("unable to save the products into the test db. %s", err.Error())
	}

	fakeId := 150
	if err = utils.LoadProductsIntoStore(int64(fakeId), db, "../candy_data.csv"); err == nil {
		t.Errorf("store with id %d does not exist", fakeId)
	}

	if err = utils.LoadProductsIntoStore(int64(id), db, "incorrect path"); err == nil {
		t.Errorf("incorrect path to a csv was given. err: %s", err.Error())
	}
}