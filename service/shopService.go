package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/terminal/location"
	"github.com/stripe/stripe-go/v75/terminal/reader"
)

var db *sql.DB

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

func SaveReader(params *stripe.TerminalReaderParams, readerStore *store.ReaderStore) error {

	reader, err := reader.New(params)
	if err != nil {
		return fmt.Errorf("unable to create a new reader, error: %s", err.Error())
	}

	storedReader := models.Reader{
		Id:         reader.ID,
		Name:       reader.Label,
		LocationId: reader.Location.ID,
	}

	_, err = readerStore.Save(&storedReader)

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
