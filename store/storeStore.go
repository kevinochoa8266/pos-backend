package store

import (
	"database/sql"

	"github.com/kevinochoa8266/pos-backend/models"
)

//var storeNotFound = errors.New("The store with this id can not be found.")

type storeStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *storeStore {
	return &storeStore{db: db}
}

func (Store *storeStore) Save(store *models.Store) (int64, error) {
	query := `INSERT INTO store (
				name,
				address
				)
				VALUES(?, ?);
	`
	result, err := Store.db.Exec(query, store.Name, store.Address)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
