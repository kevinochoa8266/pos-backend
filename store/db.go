package store

import (
	"database/sql"

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
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		address TEXT
		);
		`)
	if storeErr != nil {
		return storeErr
	}
	_, empErr := db.Exec(`CREATE TABLE IF NOT EXISTS employee(
		id INTEGER PRIMARY KEY,
		fullName TEXT NOT NULL,
		phoneNumber TEXT NOT NULL,
		address TEXT,
		storeId INTEGER,
		FOREIGN KEY (storeId) REFERENCES store (id)
		);
		`)
	if empErr != nil {
		return empErr
	}

	_, productErr := db.Exec(`CREATE TABLE IF NOT EXISTS product(
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		inventory INTEGER,
		storeId INTEGER,
		FOREIGN KEY(storeId) REFERENCES store(id)
		);
		`)
	if productErr != nil {
		return productErr
	}

	_, bulkErr := db.Exec(`CREATE TABLE IF NOT EXISTS bulk(
		productId TEXT,
		individualPrice INTEGER NOT NULL,
		bulkQuantity INTEGER,
		FOREIGN KEY (productId) REFERENCES product(id)
		);
		`)
	if bulkErr != nil {
		return bulkErr
	}

	_, taxErr := db.Exec(`CREATE TABLE IF NOT EXISTS tax(
		productId INTEGER,
		taxRate REAL NOT NULL,
		taxIncluded TEXT NOT NULL,
		FOREIGN KEY (productId) REFERENCES product (id)
		);
		`)
	if taxErr != nil {
		return taxErr
	}

	_, favErr := db.Exec(`CREATE TABLE IF NOT EXISTS favorite(
		productId INTEGER,
		imageUrl TEXT,
		FOREIGN KEY (productId) REFERENCES product (id)
		);
		`)
	if favErr != nil {
		return favErr
	}

	_, custErr := db.Exec(`CREATE TABLE IF NOT EXISTS customer(
		id INTEGER PRIMARY KEY,
		fullName TEXT NOT NULL,
		phoneNumber TEXT NOT NULL,
		email TEXT
		);
		`)
	if custErr != nil {
		return custErr
	}

	_, orderErr := db.Exec(`CREATE TABLE IF NOT EXISTS orders(
		id INTEGER PRIMARY KEY,
		productId INTEGER,
		customerId INTEGER,
		date TEXT NOT NULL,
		quantity INTEGER,
		FOREIGN KEY (productId) REFERENCES product (id),
		FOREIGN KEY (customerId) REFERENCES customer (id)
		);
		`)
	if orderErr != nil {
		return orderErr
	}

	return nil
}
