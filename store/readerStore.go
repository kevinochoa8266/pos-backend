package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kevinochoa8266/pos-backend/models"
)

type ReaderStore struct {
	db *sql.DB
}

func NewReaderStore(db *sql.DB) *ReaderStore {
	return &ReaderStore{db: db}
}

func (Reader *ReaderStore) Save(reader *models.Reader) (string, error) {
	query := `INSERT INTO reader (
				id,
				name,
				locationId
				)
				VALUES(?, ?, ?);
	`
	result, err := Reader.db.Exec(query, &reader.Id, &reader.Name, &reader.LocationId)
	if err != nil {
		return "", fmt.Errorf("query failed to save reader to database, errors: %s", err.Error())
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		return "", errors.New("failed to insert reader into database")
	}
	return reader.Id, nil
}

func (Reader *ReaderStore) GetAll() ([]models.Reader, error) {
	query := `SELECT * FROM reader;`

	result, err := Reader.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("query failed to fetch all readers from the database, error: %s", err.Error())
	}

	defer result.Close()

	readers := []models.Reader{}

	for result.Next() {
		reader := models.Reader{}
		err := result.Scan(&reader.Id, &reader.Name, &reader.LocationId)

		if err != nil {
			return nil, fmt.Errorf("unable to parse a row, err: %s", err.Error())
		}
		readers = append(readers, reader)
	}
	return readers, nil
}
