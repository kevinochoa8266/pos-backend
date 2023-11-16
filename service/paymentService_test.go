package service_test

import (
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/kevinochoa8266/pos-backend/utils"
	"github.com/stripe/stripe-go/v75"
)

var db, _ = store.GetConnection(":memory:")
var productStore = store.NewProductStore(db)
var customerStore = store.NewCustomerStore(db)
var orderStore = store.NewOrderStore(db)
var products = []models.ProductOrder{}

var coke250_inventory = 113
var diabolin_inventory = 24

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	stripe.Key = os.Getenv("STRIPE_API_KEY")

	err = store.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	shopStore := store.NewShopStore(db)
	id, err := shopStore.Save(&models.Store{Id: "FF", Name: "testStore", Address: "123 abc", City: "miami",
		State: "FL", Country: "USA", Postal: "33177"})
	if err != nil {
		panic(err)
	}

	if err = utils.LoadProductsIntoStore(id, db, "../candy_data.csv"); err != nil {
		panic(err)
	}

	products = append(products, models.ProductOrder{ProductId: "2", Quantity: 3, Price: 2500, BoughtInBulk: false})
	products = append(products, models.ProductOrder{ProductId: "7023", Quantity: 1, Price: 8900, BoughtInBulk: true})
}

func TestSaveOrder(t *testing.T) {
	customerId, err := customerStore.Save(
		&models.Customer{
			Id:          "cu-123",
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "305-687-4999",
			Email:       "john.doe@gmail.com",
			Address:     "123 AVE",
		})

	if err != nil {
		t.Error("unable to save a customer in the test database for testing")
	}

	var payment = models.Payment{
		OrderTotal: 11070,
		Products:   products,
		CustomerEmail: "john.doe@gmail.com",
		ReaderId:   "reader123",
	}

	date := time.Now().Unix()

	paymentId := "payment123"

	err = service.SaveOrder(paymentId, date, payment, orderStore, customerId)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %s", err)
	}
}

func TestProcessInventory(t *testing.T) {
	payment := models.Payment{
		OrderTotal: 11070,
		Products:   products,
		CustomerEmail: "john.doe@gmail.com",
		ReaderId:   "reader123",
	}

	err := service.ProcessInventory(payment, productStore)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %s", err)
	}

	product1, err := productStore.Get(products[0].ProductId)
	if err != nil {
		t.Fatalf("Error getting product %s: %v", product1.Id, err)
	}

	var expected_coke_inventory = coke250_inventory - products[0].Quantity

	if expected_coke_inventory != product1.Inventory {
		t.Errorf("Incorrect inventory for product %s. Expected: %d, Actual: %d", product1.Id, 110, product1.Inventory)
	}

	product2, err := productStore.Get(products[1].ProductId)
	if err != nil {
		t.Fatalf("Error getting product %s: %v", product2.Id, err)
	}

	var expected_diabolin_inventory = diabolin_inventory - (products[1].Quantity * product2.ItemsInPacket)

	if expected_diabolin_inventory != product2.Inventory {
		t.Errorf("Incorrect inventory for product %s. Expected: %d, Actual: %d", product2.Id, expected_diabolin_inventory, product2.Inventory)
	}

}
