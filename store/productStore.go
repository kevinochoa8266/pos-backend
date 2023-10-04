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
			bulkPrice,
			inventory,
			storeId
			)
			VALUES(?,?,?,?,?)`
	result, err := ps.db.Exec(query, &product.Id, &product.Name, &product.BulkPrice, &product.Inventory, &product.StoreId)
	if err != nil {
		return "", fmt.Errorf("error occurred saving %s into the database, err: %s", product.Name, err.Error())
	}
	affectedRows, err := result.RowsAffected()

	if err != nil || affectedRows != 1 {
		return "", fmt.Errorf("could not insert product with id %s into the database", product.Id)
	}

	if product.ItemsInPacket != 0 {
		err := ps.AddIndividualPrice(product.Id, product.UnitPrice, product.ItemsInPacket)
		if err != nil {
			return "", fmt.Errorf("%s was saved, but individual price was not. %s", product.Name, err.Error())
		}
	}

	return product.Id, nil
}

func (ps *ProductStore) Get(id string) (*models.Product, error) {
	product := models.Product{}

	query := `SELECT p.id, p.name, b.unitPrice, p.bulkPrice, p.inventory, b.itemsInPacket, p.storeId FROM product p 
			LEFT JOIN bulk b ON p.id = b.productId
			WHERE p.id = ?
			;`
	row := ps.db.QueryRow(query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var unitPrice sql.NullInt64
	var itemsInPacket sql.NullInt16
	err := row.Scan(&product.Id, &product.Name, &unitPrice, &product.BulkPrice, &product.Inventory, &itemsInPacket, &product.StoreId)
	if err != nil {
		return nil, err
	}

	if unitPrice.Valid {
		product.UnitPrice = unitPrice.Int64
		product.ItemsInPacket = int(itemsInPacket.Int16)
	}

	return &product, nil
}

func (ps *ProductStore) GetAll() ([]models.Product, error) {
	products := []models.Product{}

	query := `SELECT p.id, p.name, b.unitPrice, p.bulkPrice, p.inventory, b.itemsInPacket, p.storeId FROM product p 
	LEFT JOIN bulk b ON p.id = b.productId`

	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := models.Product{}

		var unitPrice sql.NullInt64
		var itemsInPacket sql.NullInt16

		err := rows.Scan(&product.Id, &product.Name, &unitPrice, &product.BulkPrice, &product.Inventory, &itemsInPacket, &product.StoreId)
		if err != nil {
			return nil, err
		}
		if unitPrice.Valid {
			product.UnitPrice = unitPrice.Int64
			product.ItemsInPacket = int(itemsInPacket.Int16)
		}

		products = append(products, product)
	}
	return products, nil
}

func (ps *ProductStore) Update(product *models.Product) error {
	query := `
		UPDATE product 
		SET storeId = ?, name = ?, bulkPrice = ?, inventory = ?  
		WHERE id = ?
	`
	result, err := ps.db.Exec(query, product.StoreId, product.Name, product.UnitPrice,
		product.Inventory, product.Id)
	if err != nil {
		return fmt.Errorf("unable to process query to update product, err: %s", err.Error())
	}
	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows != 1 {
		return errors.New("unable to update the given product")
	}
	if product.UnitPrice != 0 {
		query = `UPDATE bulk SET unitPrice = ?, itemsInPacket = ? WHERE productId = ?`
		if result, err = ps.db.Exec(query, product.BulkPrice, product.ItemsInPacket, product.Id); err != nil {
			return fmt.Errorf("could not update the bulk for id: %s due to, %s", product.Id, err.Error())
		}
		if affectedRows, _ = result.RowsAffected(); affectedRows != 1 {
			return fmt.Errorf("product %s's bulk row was not updated", product.Id)
		}
	}

	return nil
}

func (ps *ProductStore) Delete(product *models.Product) error {

	//check if there is a bulk row for this product.
	if product.UnitPrice != 0 {
		if err := ps.DeleteBulkRow(product.Id); err != nil {
			return err
		}
	}
	// check if there is an image attached to it.
	is := NewImageStore(ps.db)
	image, _ := is.Get(product.Id)
	if image != nil {
		if _, err := is.Delete(image.Id); err != nil {
			return fmt.Errorf("unable to delete image attached to the product with id %s, err: %s", image.Id, err.Error())
		}
	}

	query := "DELETE FROM product WHERE id = ?"

	result, err := ps.db.Exec(query, product.Id)
	if err != nil {
		return fmt.Errorf("could not delete product with id: %s: %s", product.Id, err.Error())
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil || rowsDeleted != 1 {
		return errors.New("product with the given id does not exist")
	}
	return nil
}

func (ps *ProductStore) DeleteBulkRow(id string) error {
	query := `DELETE FROM bulk WHERE productId = ?`

	result, err := ps.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("unable to delete row in bulk table with id: %s. %s", id, err.Error())
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected != 1 {
		return fmt.Errorf("there was no row with id: %s in bulk table", id)
	}
	return nil
}

func (ps *ProductStore) AddIndividualPrice(id string, unitPrice int64, bulkQuantity int) error {
	query := "INSERT INTO bulk (productId, unitPrice, itemsInPacket) VALUES (?,?,?)"

	result, err := ps.db.Exec(query, id, unitPrice, bulkQuantity)
	if err != nil {
		return fmt.Errorf("could not insert into bulk table with id: %s. %s", id, err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return fmt.Errorf("unable to insert data with id %s into the database", id)
	}
	return nil
}
