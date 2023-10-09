package service

import (
	"database/sql"
	"os"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/terminal/location"
)

func CreateLocation() (*stripe.TerminalLocation, error) {
	reader_location := models.Store{
		Address: os.Getenv("STORE_ADDRESS"),
		City:    os.Getenv("STORE_CITY"),
		State:   os.Getenv("STORE_STATE"),
		Country: os.Getenv("STORE_COUNTRY"),
		Postal:  os.Getenv("STORE_POSTAL"),
		Name:    os.Getenv("STORE_NAME"),
	}
	params := &stripe.TerminalLocationParams{
		Address: &stripe.AddressParams{
			Line1:      stripe.String(reader_location.Address),
			City:       stripe.String(reader_location.City),
			State:      stripe.String(reader_location.State),
			Country:    stripe.String(reader_location.Country),
			PostalCode: stripe.String(reader_location.Postal),
		},
		DisplayName: stripe.String(reader_location.Name),
	}
	l, err := location.New(params)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func SaveReader(readerId string, locationId string, name string, db *sql.DB) error {
	newStore := store.NewReaderStore(db)

	reader := models.Reader{
		Id:         readerId,
		Name:       name,
		LocationId: locationId,
	}

	_, err := newStore.Save(&reader)

	if err != nil {
		return err
	}

	return nil
}

func InitializeShop(shop *store.ShopStore) error {
	location, err := CreateLocation()
	if err != nil {
		return err
	}

	candyShop := models.Store{
		Id:      location.ID,
		Address: location.Address.Line1,
		City:    location.Address.City,
		State:   location.Address.State,
		Country: location.Address.Country,
		Postal:  location.Address.PostalCode,
		Name:    location.DisplayName,
	}

	_, err = shop.Save(&candyShop)
	if err != nil {
		return err
	}
	return nil
}
