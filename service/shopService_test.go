package service_test

import (
	"os"
	"testing"

	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/kevinochoa8266/pos-backend/store"
)

func TestInitializeShop(t *testing.T) {
	os.Setenv("STORE_ADDRESS", "123 Main St")
	os.Setenv("STORE_CITY", "Gville")
	os.Setenv("STORE_STATE", "CA")
	os.Setenv("STORE_COUNTRY", "US")
	os.Setenv("STORE_POSTAL", "12345")
	os.Setenv("STORE_NAME", "Test Store")

	defer func() {
		os.Unsetenv("STORE_ADDRESS")
		os.Unsetenv("STORE_CITY")
		os.Unsetenv("STORE_STATE")
		os.Unsetenv("STORE_COUNTRY")
		os.Unsetenv("STORE_POSTAL")
		os.Unsetenv("STORE_NAME")
	}()
	db, _ := store.GetConnection(":memory:")
	store.CreateSchema(db)
	shopStore := store.NewStore(db)
	err := service.InitializeShop(&shopStore)
	if err != nil {
		panic(err)
	}
	shops, err := shopStore.GetAll()
	if err != nil {
		t.Errorf("No shops found")
	}

	for i, shop := range shops {
		t.Logf("Shop %d - ID: %s, Name: %s, Address: %s", i+1, shop.Id, shop.Name, shop.Address)
	}
}
