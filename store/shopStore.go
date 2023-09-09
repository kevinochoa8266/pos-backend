package store

import (
	"database/sql"
	"errors"

	"github.com/kevinochoa8266/pos-backend/models"
)

var storeNotFound = errors.New("The store with this id can not be found.")

type shopStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *shopStore {
	return &shopStore{db: db}
}

func (Store *shopStore) Save(store *models.Store) (int64, error) {
	query := `INSERT INTO store (
				Id,
				name,
				address
				)
				VALUES(?, ?, ?);
	`
	result, err := Store.db.Exec(query, &store.Id, &store.Name, &store.Address)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (Store *shopStore) Get(id int) (*models.Store, error) {
	query := `SELECT * FROM store s WHERE s.id = ?;`

	result := Store.db.QueryRow(query, id)

	if result.Err() != nil {
		return nil, storeNotFound
	}

	shop := models.Store{}

	err := result.Scan(&shop.Id, &shop.Name, &shop.Address)
	if err != nil {
		return nil, err
	}
	return &shop, nil
}
