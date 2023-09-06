package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kevinochoa8266/pos-backend/models"
)

type ProductStore struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{db: db}
}

func (ps *ProductStore) Save(product *models.Product) (string, error) {
	query := `INSERT INTO product (
			id,
			name,
			price,
			inventory,
			storeId
			)
			VALUES(?,?,?,?,?)`
	result, err := ps.db.Exec(query, &product.Id, &product.Name, &product.Price, &product.Inventory, &product.StoreId)
	if err != nil {
		return "", err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows != 1 {
		return "", fmt.Errorf("could not insert product with id %s into the database", product.Id)
	}

	if product.TaxRate != 0.00 {
		err := ps.addTax(product.Id, product.TaxIncluded, product.TaxRate)
		if err != nil {
			return "", fmt.Errorf("product was saved, but the tax was not added. %s", err.Error())
		}
	}

	if product.Bul

	return product.Id, nil
}

func (ps *ProductStore) Get(id string) (*models.Product, error) {
	product := models.Product{}

	query := `SELECT * FROM product p WHERE p.id = ?`
	row := ps.db.QueryRow(query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&product.Id, &product.StoreId, &product.Name, &product.Price, &product.Inventory)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (ps *ProductStore) GetAll() ([]models.Product, error) {
	products := []models.Product{}

	query := "SELECT * FROM product p"

	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := models.Product{}

		err := rows.Scan(&product.Id, &product.StoreId, &product.Name, &product.Price, &product.Inventory)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (ps *ProductStore) Update(product *models.Product) error {
	query := `
		UPDATE product 
		SET storeId = ?, name = ?, price = ?, inventory = ?
		WHERE id = ?
	`
	result, err := ps.db.Exec(query, product.StoreId, product.Name, product.Price,
		product.Inventory, product.Id)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows != 1 {
		return errors.New("unable to update the given product")
	}
	return nil
}

func (ps *ProductStore) Delete(id string) error {
	query := "DELETE FROM product WHERE id = ?"

	result, err := ps.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete product with id: %s: %s", id, err.Error())
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil || rowsDeleted != 1 {
		return errors.New("product with the given id does not exist")
	}
	return nil
}

func (ps *ProductStore) addTax(id string, taxIncluded bool, taxRate float32) (error) {
	query := "INSERT INTO tax (id, taxIncluded, taxRate) VALUES(?,?,?)"

	result, err := ps.db.Exec(query, id, taxIncluded, taxRate)
	if err != nil {
		return fmt.Errorf("query failed to add the the tax row with id: %s. %s",id, err.Error())
	}
	affectedRows, err := result.RowsAffected(); if err != nil || affectedRows != 1 {
		return fmt.Errorf("tax row was unable to be inserted with id: %s", id)
	}

	return nil
}
