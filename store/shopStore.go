package store

import (
	"database/sql"
	"fmt"

	"github.com/kevinochoa8266/pos-backend/models"
)

type ShopStore struct {
	db *sql.DB
}

func NewShopStore(db *sql.DB) *ShopStore {
	return &ShopStore{db: db}
}

func (Store *ShopStore) Save(store *models.Store) (string, error) {
	query := `INSERT INTO store (
				Id,
				address,
				city,
				state,
				country,
				postal,
				name
				)
				VALUES(?, ?, ?, ?, ?, ?, ?);
	`
	_, err := Store.db.Exec(query, &store.Id, &store.Address, &store.City, &store.State, &store.Country, &store.Postal, &store.Name)
	if err != nil {
		return "", fmt.Errorf("unable to save store into database, error: %s", err.Error())
	}

	return store.Id, nil
}

func (Store *ShopStore) Get(id string) (*models.Store, error) {
	query := `SELECT * FROM store s WHERE s.id = ?;`

	result := Store.db.QueryRow(query, id)

	if result.Err() != nil {
		return nil, fmt.Errorf("unable to fetch store with id: %s, error: %s", id, result.Err().Error())
	}

	shop := models.Store{}

	err := result.Scan(&shop.Id, &shop.Address, &shop.City, &shop.State, &shop.Country, &shop.Postal, &shop.Name)
	if err != nil {
		return nil, fmt.Errorf("unable to parse a row, err: %s", err.Error())
	}
	return &shop, nil
}

func (Store *ShopStore) GetAll() ([]models.Store, error) {
	query := `SELECT * FROM store;`

	result, err := Store.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("unable to fetch all stores from database, error: %s", err.Error())
	}

	defer result.Close()

	shops := []models.Store{}

	for result.Next() {
		shop := models.Store{}
		err := result.Scan(&shop.Id, &shop.Address, &shop.City, &shop.State, &shop.Country, &shop.Postal, &shop.Name)

		if err != nil {
			return nil, fmt.Errorf("unable to parse a row, err: %s", err.Error())
		}
		shops = append(shops, shop)
	}
	return shops, nil
}

func (Store *ShopStore) Update(store *models.Store) error {
	query := `UPDATE store SET Name = ?, Address = ? WHERE Id = ?`

	result, err := Store.db.Exec(query, &store.Id, &store.Address, &store.City, &store.State, &store.Country, &store.Postal, &store.Name)

	if err != nil {
		return fmt.Errorf("unable to update store in database with id: %s, error: %s", store.Id, err.Error())
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil || rowsUpdated != 1 {
		return err
	}
	return nil
}

// TODO: Clean up the deletes from all foreign keys before deleting this store
// func (Store  shopStore) Delete(Id int) error {
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
