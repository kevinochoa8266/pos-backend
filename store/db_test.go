package store_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/store"
)

var dbUrl = ":memory:"
var name = "John"
var number = "9417160432"
var address = "123 abc"

var DB, _ = store.GetConnection(dbUrl)

func init() {
	store.BuildTestDb(DB)
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
