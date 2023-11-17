package service_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/service"
	"github.com/kevinochoa8266/pos-backend/store"
	"github.com/kevinochoa8266/pos-backend/utils"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
	readertesthelpers "github.com/stripe/stripe-go/v75/testhelpers/terminal/reader"
)

var db, _ = store.GetConnection(":memory:")
var productStore = store.NewProductStore(db)
var customerStore = store.NewCustomerStore(db)
var orderStore = store.NewOrderStore(db)
var shopStore = store.NewShopStore(db)
var readerStore = store.NewReaderStore(db)
var products = []models.ProductOrder{}
var Req struct {
	RegistrationCode string `json:"registration_code"`
	Label            string `json:"label"`
}
var readerId = ""
var coke250_inventory = 113
var diabolin_inventory = 24

func init() {

	if _, inCI := os.LookupEnv("GITHUB_ACTIONS"); inCI {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Errorf("Error loading environment variables: %s", err))
		}
	} else {
		err := godotenv.Load("../.env")
		if err != nil {
			panic(fmt.Errorf("Error loading environment variables: %s", err))
		}
	}

	stripe.Key = os.Getenv("STRIPE_API_KEY")

	err := store.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	err = service.InitializeShop(shopStore)

	if err != nil {
		panic(err)
	}

	stores, err := shopStore.GetAll()
	if err != nil {
		panic(err)
	}

	Req.RegistrationCode = "simulated-wpe"
	Req.Label = "payment-service-testing-reader"

	id, err := service.SaveReader(Req, readerStore, shopStore)

	readerId = id

	if err != nil {
		panic(err)
	}

	if err = utils.LoadProductsIntoStore(stores[0].Id, db, "../candy_data.csv"); err != nil {
		panic(err)
	}

	products = append(products, models.ProductOrder{ProductId: "2", Quantity: 3, Price: 2500, BoughtInBulk: false})
	products = append(products, models.ProductOrder{ProductId: "7023", Quantity: 1, Price: 8900, BoughtInBulk: true})

	_, err = customerStore.Save(
		&models.Customer{
			Id:          "cu-123",
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "305-687-4999",
			Email:       "john.doe@gmail.com",
			Address:     "123 AVE",
		})

	if err != nil {
		panic(err)
	}
}

func TestSaveOrder(t *testing.T) {
	customerEmail := "john.doe@gmail.com"

	customerId, err := customerStore.GetByEmail(customerEmail)

	if err != nil {
		t.Errorf("unable to retrieve customer with email: %s, error: %s", customerEmail, err)
	}

	var payment = models.Payment{
		OrderTotal:    11070,
		Products:      products,
		CustomerEmail: customerEmail,
		ReaderId:      "reader123",
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
		OrderTotal:    11070,
		Products:      products,
		CustomerEmail: "john.doe@gmail.com",
		ReaderId:      "reader123",
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

func TestCreatePaymentIntentParams(t *testing.T) {
	// This payment model represents a customer who wants a reciept and has a customer profile.
	payment := models.Payment{
		OrderTotal:    11070,
		Products:      products,
		CustomerEmail: "john.doe@gmail.com",
		ReaderId:      "reader123",
	}

	expectedId := "cu-123"

	_, id, err := service.CreatePaymentIntentParams(payment, customerStore)

	if err != nil {
		t.Errorf("unable to create payment intent params with email: %s, err: %s", payment.CustomerEmail, err)
	}

	if id != expectedId {
		t.Errorf("incorrect id was returned, expected: %s, actual: %s", expectedId, id)
	}

	// This is the payment model that represents a customer who does not want a reciept.
	payment = models.Payment{
		OrderTotal:    11070,
		Products:      products,
		CustomerEmail: "",
		ReaderId:      "reader123",
	}

	expectedId = ""

	_, id, err = service.CreatePaymentIntentParams(payment, customerStore)

	if err != nil {
		t.Errorf("unable to create payment intent params with email: %s, err: %s", payment.CustomerEmail, err)
	}

	if id != expectedId {
		t.Errorf("incorrect id was returned, expected: %s, actual: %s", expectedId, id)
	}

	// This is the payment model that represents a customer who wants a reciept but does not have a customer profile.
	payment = models.Payment{
		OrderTotal:    11070,
		Products:      products,
		CustomerEmail: "jane.doe@gmail.com",
		ReaderId:      "reader123",
	}

	expectedId = ""

	_, id, err = service.CreatePaymentIntentParams(payment, customerStore)

	if err != nil {
		t.Errorf("unable to create payment intent params with email: %s, err: %s", payment.CustomerEmail, err)
	}

	if id != expectedId {
		t.Errorf("incorrect id was returned, expected: %s, actual: %s", expectedId, id)
	}

}

func TestProcessAndSimulatePayment(t *testing.T) {

	payment := models.Payment{
		OrderTotal:    11070,
		Products:      products,
		CustomerEmail: "",
		ReaderId:      readerId,
	}

	params, _, err := service.CreatePaymentIntentParams(payment, customerStore)

	if err != nil {
		t.Errorf("unable to create payment intent params with email: %s, err: %s", payment.CustomerEmail, err)
	}

	pi, err := paymentintent.New(params)

	if err != nil {
		t.Errorf("unable to create a new payment intent, error: %s", err.Error())
	}

	err = service.ProcessPayment(readerId, pi.ID)

	if err != nil {
		t.Errorf("unable to process payment with id: %s, error: %s", pi.ID, err)
	}

	presentPaymentParams := &stripe.TestHelpersTerminalReaderPresentPaymentMethodParams{}
	resp, err := readertesthelpers.PresentPaymentMethod(readerId, presentPaymentParams)

	if err != nil {
		t.Errorf("the reader: %s was unable to simulate the payment, error: %s", readerId, err)
	}

	if resp.Action.Status == "failed" {
		t.Errorf("unable to simulate payment: %s", resp.Action.FailureMessage)
	}
	
}

func TestTransactionProcess(t *testing.T) {

	payment := models.Payment{
		OrderTotal:    11070,
		Products:      products,
		CustomerEmail: "",
		ReaderId:      readerId,
	}

	ExpectedResponseStatus := "succeeded"

	response, err := service.TransactionProcess(payment, orderStore, productStore, customerStore)

	if err != nil {
		t.Error("The transaction was unable to process", err)
	}

	if ExpectedResponseStatus != response {
		t.Errorf("transaction failed and was supposed to be successful")
	}

}
