package store

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/kevinochoa8266/pos-backend/models"
	_ "modernc.org/sqlite"
)

// Passes in the url, gets the connection, and checks if it is valid with a Ping.
func GetConnection(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbUrl)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseConnection(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

func CreateSchema(db *sql.DB) error {
	if _, err := db.Exec("PRAGMA foreign_keys = ON", nil); err != nil {
		return err
	}
	_, storeErr := db.Exec(`CREATE TABLE IF NOT EXISTS store(
		id TEXT PRIMARY KEY,
		address TEXT NOT NULL,
		city TEXT NOT NULL,
		state TEXT NOT NULL,
		country TEXT NOT NULL,
		postal TEXT NOT NULL,
		name TEXT NOT NULL
		);
		`)
	if storeErr != nil {
		return storeErr
	}
	_, readerErr := db.Exec(`CREATE TABLE IF NOT EXISTS reader(
		id TEXT PRIMARY KEY,
		name TEXT,
		locationId TEXT NOT NULL,
		FOREIGN KEY (locationId) REFERENCES store (id) 
		);
		`)
	if readerErr != nil {
		return readerErr
	}
	_, empErr := db.Exec(`CREATE TABLE IF NOT EXISTS employee(
		id INTEGER PRIMARY KEY,
		fullName TEXT NOT NULL,
		phoneNumber TEXT,
		address TEXT NOT NULL,
		storeId INTEGER NOT NULL,
		FOREIGN KEY (storeId) REFERENCES store (id)
		);
		`)
	if empErr != nil {
		return empErr
	}

	_, productErr := db.Exec(`CREATE TABLE IF NOT EXISTS product(
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		bulkPrice INTEGER NOT NULL,
		inventory INTEGER NOT NULL,
		storeId INTEGER NOT NULL,
		FOREIGN KEY(storeId) REFERENCES store(id)
		);
		`)
	if productErr != nil {
		return productErr
	}

	_, bulkErr := db.Exec(`CREATE TABLE IF NOT EXISTS bulk(
		productId TEXT PRIMARY KEY,
		unitPrice INTEGER NOT NULL,
		itemsInPacket INTEGER NOT NULL,
		FOREIGN KEY (productId) REFERENCES product(id)
		);
		`)
	if bulkErr != nil {
		return bulkErr
	}

	_, favErr := db.Exec(`CREATE TABLE IF NOT EXISTS favorite(
		id TEXT PRIMARY KEY,
		data BLOB NOT NULL,
		FOREIGN KEY (id) REFERENCES product (id)
		);
		`)
	if favErr != nil {
		return favErr
	}

	_, custErr := db.Exec(`CREATE TABLE IF NOT EXISTS customer(
		id TEXT PRIMARY KEY,
		firstName TEXT NOT NULL,
		lastName TEXT NOT NULL,
		phoneNumber TEXT,
		email TEXT UNIQUE,
		address TEXT 
		);
		`)
	if custErr != nil {
		return custErr
	}

	_, orderErr := db.Exec(`CREATE TABLE IF NOT EXISTS orders(
		id TEXT NOT NULL,
		date DATE NOT NULL,
		quantity INTEGER NOT NULL,
		priceAtPurchase INTEGER NOT NULL,
		productId INTEGER NOT NULL,
		customerId TEXT,
		FOREIGN KEY (productId) REFERENCES product (id),
		FOREIGN KEY (customerId) REFERENCES customer (id)
		);
		`)
	if orderErr != nil {
		return orderErr
	}

	return nil
}

func BuildTestDb(db *sql.DB) {
	var name = "John"
	var number = "9417160432"
	var address = "123 abc"

	var storeId = "FF"
	var storeAddress = "123 abc street"
	var storeCity = "Gainesville"
	var storeState = "FL"
	var storeCountry = "USA"
	var storePostal = "050037"
	var storeName = "XYZ Store"

	var productName = "Chocolate"
	var price int64 = 5
	var inventory = 100

	if err := CreateSchema(db); err != nil {
		panic(err)
	}
	shopStore := NewShopStore(db)
	_, err := shopStore.Save(&models.Store{
		Id: storeId,
		Name: storeName,
		Address: storeAddress,
		City: storeCity,
		State: storeState,
		Country: storeCountry,
		Postal: storePostal,
	}); if err != nil {
		panic(err)
	}
	es := NewEmployeeStore(db)

	for i := 0; i < 3; i++ {
		if _, err := es.Save(&models.Employee{
			FullName: name,
			Phone: number,
			Address: address,
			StoreId: storeId,
		}); err != nil {
			panic(err)
		}
	}
	ps := NewProductStore(db)

	// add products
	for i := 0; i < 10; i++ {
		_, err := ps.Save(&models.Product{
			Id:            strconv.Itoa(i),
			Name:          productName,
			UnitPrice:     price,
			Inventory:     inventory,
			BulkPrice:     price * 5,
			ItemsInPacket: 10,
			StoreId:       storeId})
		if err != nil {
			panic(err)
		}
	}

	//add orders
	order := models.Order{
		Id:              "abr_123",
		Date:            time.Now(),
		Quantity:        3,
		ProductId:       "",
		PriceAtPurchase: 320,
		CustomerId:      0,
	}
	orderStore := NewOrderStore(db)
	for i := 0; i < 4; i++ {
		productId := strconv.Itoa(i)
		product, _ := ps.Get(productId)
		order.ProductId = productId
		order.PriceAtPurchase = product.BulkPrice
		err := orderStore.Save(&order)
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < 3; i++ {
		order.Id = uuid.NewString()
		if err := orderStore.Save(&order); err != nil {
			panic(err)
		}
	}
}
