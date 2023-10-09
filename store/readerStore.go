package store

import (
	"database/sql"

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
	_, err := Reader.db.Exec(query, &reader.Id, &reader.Name, &reader.LocationId)
	if err != nil {
		return "", err
	}

	return reader.Id, nil
}

func (Reader *ReaderStore) GetAll() ([]models.Reader, error) {
	query := `SELECT * FROM reader;`

	result, err := Reader.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	readers := []models.Reader{}

	for result.Next() {
		reader := models.Reader{}
		err := result.Scan(&reader.Id, &reader.Name, &reader.LocationId)

		if err != nil {
			return nil, err
		}
		readers = append(readers, reader)
	}
	return readers, nil
}
