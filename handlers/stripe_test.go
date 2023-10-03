package handlers_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/kevinochoa8266/pos-backend/handlers"
	"github.com/stretchr/testify/assert"
)

func TestStripeCreateLocation(t *testing.T)  {
	os.Setenv("STORE_ADDRESS", "123 Main St")
	os.Setenv("STORE_CITY", "Anytown")
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

	locationId, err := handlers.CreateLocation()

	fmt.Println(locationId)

	assert.NoError(t, err, "CreateLocation should not return an error")
}
