package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

func TestReaderSave(t *testing.T) {
	newStore := store.NewReaderStore(db)

	reader := models.Reader{
		Id:         "XZY",
		Name:       "reader-one",
		LocationId: "FF",
	}

	_, err := newStore.Save(&reader)
	if err != nil {
		t.Error("could not save a new store in the database")
	}
}

func TestReaderGetAll(t *testing.T) {
	newStore := store.NewReaderStore(db)
	reader := models.Reader{
		Id:         "abc",
		Name:       "reader-one",
		LocationId: "FF",
	}

	_, err := newStore.Save(&reader)
	if err != nil {
		t.Error("could not save a new store in the database")
	}

	readers, err := newStore.GetAll()

	if err != nil {
		t.Error("Failed to fetch all of the stores from the database.")
	}

	if len(readers) == 0 {
		t.Error("No stores were retrieved")
	}
}
