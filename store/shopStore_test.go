package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var shop = models.Store{
	Id:      6,
	Name:    "Candy Store",
	Address: "Medellin",
}

func TestStoreSave(t *testing.T) {
	newStore := store.NewStore(db)

	store := models.Store{
		Name:    "ABC Store",
		Address: "345 AVE",
	}

	_, err := newStore.Save(&store)
	if err != nil {
		t.Error("could not save a new store in the database")
	}
}

func TestStoreGet(t *testing.T) {
	shopStore := store.NewStore(db)

	shopStore.Save(&shop)

	id := 6
	_, err := shopStore.Get(id)

	if err != nil {
		t.Errorf("The store with id: %d was not found.", id)
	}

	id2 := 100
	_, err = shopStore.Get(id2)

	if err == nil {
		t.Errorf("was not supposed to retrieve a store with the id: %d", id2)
	}

}
