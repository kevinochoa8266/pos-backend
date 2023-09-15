package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

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

	id := 1
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

func TestStoreGetAll(t *testing.T) {
	shopStore := store.NewStore(db)

	shops, err := shopStore.GetAll()

	if err != nil {
		t.Error("Failed to fetch all of the stores from the database.")
	}

	if len(shops) == 0 {
		t.Error("No stores were retrieved")
	}
}

func TestStoreUpdate(t *testing.T) {
	shopStore := store.NewStore(db)

	shop, err := shopStore.Get(3)

	if err != nil {
		t.Error("Shop was not retrieved.")
	}

	shop.Name = "Starbucks"

	err = shopStore.Update(shop)

	if err != nil {
		t.Errorf("could not update store with id: %d", shop.Id)
	}
}
// TODO: This needs to also remove all rows that have a foreign key that reference the given store.
// func TestStoreDelete(t *testing.T) {
// 	shopStore := store.NewStore(db)

// 	id := 1

// 	err := shopStore.Delete(id)

// 	if err != nil {
// 		t.Errorf("could not delete store with id: %d", id)
// 	}

// 	id = 20

// 	err = shopStore.Delete(id)

// 	if err == nil {
// 		t.Errorf("store with the following Id should not exist: %d", id)
// 	}
// }
