package store

import (
	"database/sql"
	"errors"

	"github.com/kevinochoa8266/pos-backend/models"
)

type ProductStore struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{db: db}
}

func (ps *ProductStore) Save(product *models.Product) (int64, error) {
	query := `INSERT INTO product (
			name,
			price,
			inventory,
			storeId
			)
			VALUES(?,?,?,?)`
	result, err := ps.db.Exec(query, &product.Name, &product.Price, &product.Inventory, &product.StoreId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ps *ProductStore) Get(id int) (*models.Product, error) {
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

func (ps *ProductStore) Update(*models.Product) error {
	return errors.ErrUnsupported
}

func (ps *ProductStore) Delete(*models.Product) error {
	return errors.ErrUnsupported
}
