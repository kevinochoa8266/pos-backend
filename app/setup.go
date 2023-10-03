package app

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/kevinochoa8266/pos-backend/utils"
)

func SetupApp() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	dbUrl := os.Getenv("DB_URL")
	// Create the schema to the database
	db, err := store.GetConnection(dbUrl)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(3)
	defer db.Close()

	if err = store.CreateSchema(db); err != nil {
		return err
	}

	// Check if a store exists.
	shopStore := store.NewStore(db)
	stores, err := shopStore.GetAll()
	if err != nil {
		return err
	}
	if len(stores) == 0 {
		name := os.Getenv("STORE_NAME")
		address := os.Getenv("STORE_ADDRESS")
		shop := models.Store{
			Name:    name,
			Address: address,
		}

		_, err := shopStore.Save(&shop)
		if err != nil {
			return err
		}
	}
	productStore := store.NewProductStore(db)

	products, err := productStore.GetAll()
	if err != nil {
		return err
	}

	if len(products) == 0 {
		stores, err = shopStore.GetAll()
		if err != nil {
			return err
		}
		storeId := stores[0].Id
		if err = utils.LoadProductsIntoStore(int64(storeId), db, "candy_data.csv"); err != nil {
			return err
		}
	}
	return nil
}