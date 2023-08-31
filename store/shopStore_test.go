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
