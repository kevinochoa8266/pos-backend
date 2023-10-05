package service

import (
	"os"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/terminal/location"
	"github.com/stripe/stripe-go/v75/terminal/reader"
)

func CreateLocation() (string, error) {
	reader_location := models.Reader{
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
		return "", err
	}
	return l.ID, nil
}

func RegisterReader(locationId string) (string, error) {

	params := &stripe.TerminalReaderParams{
		Location:         stripe.String(locationId),
		RegistrationCode: stripe.String("simulated-wpe"),
	}

	reader, err := reader.New(params)
	if err != nil {
		return "", err
	}

	return reader.ID, nil
}

func InitializeShop(shop *store.ShopStore) error {
	locationId, err := CreateLocation()
	if err != nil {
		return err
	}

	readerId, err := RegisterReader(locationId)
	if err != nil {
		return err
	}

	name := os.Getenv("STORE_NAME")
	address := os.Getenv("STORE_ADDRESS")
	candyShop := models.Store{
		Id:       locationId,
		Name:     name,
		Address:  address,
		ReaderId: readerId,
	}

	_, err = shop.Save(&candyShop)
	if err != nil {
		return err
	}
	return nil
}
