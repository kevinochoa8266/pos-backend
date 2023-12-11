package utils_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/kevinochoa8266/pos-backend/utils"
)

func TestLoadProducts(t *testing.T) {
	t.SkipNow()
	db, _ := store.GetConnection(":memory:")

	store.CreateSchema(db)
	shopStore := store.NewShopStore(db)
	id, err := shopStore.Save(&models.Store{Id: "FF", Name: "testStore", Address: "123 abc", City: "miami",
		State: "FL", Country: "USA", Postal: "33177"})
	if err != nil {
		t.Errorf("unable to set up db for testing.")
	}

	if err = utils.LoadProductsIntoStore(id, db, "../candy_data.csv"); err != nil {
		t.Log(id)
		t.Errorf("unable to save the products into the test db. %s", err.Error())
	}

	fakeId := "XYZ"
	if err = utils.LoadProductsIntoStore(fakeId, db, "../candy_data.csv"); err == nil {
		t.Errorf("store with id %s does not exist", fakeId)
	}

	if err = utils.LoadProductsIntoStore(id, db, "incorrect path"); err == nil {
		t.Errorf("incorrect path to a csv was given. err: %s", err.Error())
	}
}
