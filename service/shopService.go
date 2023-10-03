package service

import (
	"os"

	"github.com/kevinochoa8266/pos-backend/handlers"
	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

func InitializeShop(shop *store.ShopStore) error {
	locationId, err := handlers.CreateLocation()
	if err != nil {
		return err
	}

	name := os.Getenv("STORE_NAME")
	address := os.Getenv("STORE_ADDRESS")
	candyShop := models.Store{
		Id:      locationId,
		Name:    name,
		Address: address,
	}

	_, err = shop.Save(&candyShop)
	if err != nil {
		return err
	}
	return nil
}
