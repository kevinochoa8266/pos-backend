package store

import (
	"database/sql"
	"errors"

	"github.com/kevinochoa8266/pos-backend/models"
)

var errFoo = errors.New("the shop with this id can not be found")

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
		return nil, errFoo
	}

	shop := models.Store{}

	err := result.Scan(&shop.Id, &shop.Name, &shop.Address)
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

func (Store *shopStore) GetAll() ([]models.Store, error) {
	query := `SELECT * FROM store;`

	result, err := Store.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	shops := []models.Store{}

	for result.Next() {
		shop := models.Store{}
		err := result.Scan(&shop.Id, &shop.Name, &shop.Address)

		if err != nil {
			return nil, err
		}
		shops = append(shops, shop)
	}
	return shops, nil
}

func (Store *shopStore) Update(store *models.Store) error {
	query := `UPDATE store SET Name = ?, Address = ? WHERE Id = ?`

	result, err := Store.db.Exec(query, &store.Name, &store.Address, &store.Id)

	if err != nil {
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil || rowsUpdated != 1 {
		return err
	}
	return nil
}

// TODO: Clean up the deletes from all foreign keys before deleting this store
// func (Store *shopStore) Delete(Id int) error {
// 	query := `DELETE FROM store WHERE Id = ?`

// 	result, err := Store.db.Exec(query, Id)

// 	if err != nil {
// 		return err
// 	}

// 	rowsDeleted, err := result.RowsAffected()
// 	if err != nil || rowsDeleted != 1 {
// 		return errors.New("shop with the given id does not exist")
// 	}
// 	return nil
// }
