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

	if len(image.Data) == 0 {
		return fmt.Errorf("can not pass in an empty image to the database")
	}
	images, err := is.GetAll()
	if err != nil {
		return err
	}
	if len(images) > 10 {
		return fmt.Errorf("user has reached their images insert length")
	}
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

func (is *ImageStore) GetAll() ([]models.Image, error) {
	query := "SELECT id, data from favorite"

	rows, err := is.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed to get all images, err: %s", err.Error())
	}
	defer rows.Close()

	images := []models.Image{}
	for rows.Next() {
		image := models.Image{}
		if err := rows.Scan(&image.Id, &image.Data); err != nil {
			return nil, fmt.Errorf("unable to scan image row from the database, err: %s", err.Error())
		}
		images = append(images, image)
	}
	return images, nil
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
		return fmt.Errorf("the image was not successfully updated to the database")
	}
	return nil
}

func (is *ImageStore) Delete(id string) (string, error) {
	query := "DELETE FROM favorite where id = ?"

	result, err := is.db.Exec(query, id)
	if err != nil {
		return "", fmt.Errorf("query failed to delete image %s, err: %s", id, err.Error())
	}
	if affectedRows, _ := result.RowsAffected(); affectedRows != 1 {
		return "", fmt.Errorf("image %s was not deleted from the database", id)
	}
	return id, nil
}
