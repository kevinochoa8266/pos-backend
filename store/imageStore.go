package store

import (
	"bytes"
	"database/sql"
	"fmt"

	"github.com/kevinochoa8266/pos-backend/models"
)

type ImageStore struct {
	db *sql.DB
}

func NewImageStore(db *sql.DB) *ImageStore {
	return &ImageStore{db: db}
}

func (is *ImageStore) Save(image *models.Image) error {
	query := "INSERT INTO favorite (id, data) VALUES(?, ?)"

	result, err := is.db.Exec(query, &image.Id, &image.Data)
	if err != nil {
		return fmt.Errorf("failed to insert image into db, err: %s", err.Error())
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected != 1 {
		return fmt.Errorf("image was not inserted into the database")
	}
	return nil
}

// Gets a image from the db if the id does exist. Returns a errNoRows if no rows are matched from the query.
func (is *ImageStore) Get(id string) (*models.Image, error) {
	query := "SELECT id, data from favorite where id = ?"

	row := is.db.QueryRow(query, id)
	if row.Err() != nil {
		return nil, fmt.Errorf("query failed to get image from id: %s, err: %s", id, row.Err().Error())
	}
	image := models.Image{}
	if err := row.Scan(&image.Id, &image.Data); err != nil {
		return nil, err
	}
	return &image, nil
}

func (is *ImageStore) Update(image *models.Image) error {
	query := "UPDATE favorite SET data = ? WHERE id = ?"

	_, err := is.db.Exec(query, image.Data, image.Id)
	if err != nil {
		return fmt.Errorf("failed to execute update query, err: %s", err.Error())
	}
	updatedImage, err := is.Get(image.Id)
	if err != nil {
		return sql.ErrNoRows
	}
	if !bytes.Equal(image.Data, updatedImage.Data) {
		return fmt.Errorf("image with id %s was not updated", image.Id)
	}
	return nil
}

func (is *ImageStore) Delete(id string) (string, error) {
	query := "DELETE FROM favorite where id = ?"

	_, err := is.db.Exec(query, id)
	if err != nil {
		return "", fmt.Errorf("query failed to delete image %s, err: %s", id, err.Error())
	}
	if _, err := is.Get(id); err != sql.ErrNoRows {
		return "", fmt.Errorf("image %s was not deleted from the database", id)
	}
	return id, nil
}
