package store_test

import (
	"database/sql"
	"os"
	"strconv"
	"testing"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var imageStore = store.NewImageStore(db)
var imageData, _ = os.ReadFile("../snickers.png")

var image = models.Image{Id: "3", Data: imageData}

func TestSave(t *testing.T) {
	if err := imageStore.Save(&image); err != nil {
		t.Errorf("unable to save image for product into the database. err: %s", err.Error())
	}

	image.Id = "id does not exist"
	if err := imageStore.Save(&image); err == nil {
		t.Error("foreign key error should have thrown trying to save image into the db")
	}
}

func TestGetImage(t *testing.T) {
	image.Id = "3"
	imageStore.Save(&image)
	image, err := imageStore.Get(image.Id)
	if err != nil {
		t.Errorf("unbale to get image with id %s, err: %s", image.Id, err.Error())
	}

	_, err = imageStore.Get("id does not exist")
	if err != sql.ErrNoRows {
		t.Error("expected to not get a image from the fake id")
	}
}

func TestGetAllImages(t *testing.T) {
	ps := store.NewProductStore(db)
	for i := 20; i < 23; i++ {
		product := models.Product{}
		product.Id = strconv.Itoa(i)
		product.StoreId = 1
		if _, err := ps.Save(&product); err != nil {
			t.Error("unable to save products into db for testing getting images")
		}
	}

	for i := 20; i < 23; i++ {
		image.Id = strconv.Itoa(i)
		err := imageStore.Save(&image)
		if err != nil {
			t.Error("unable to store image into the db")
		}
	}

	images, err := imageStore.GetAll()
	if err != nil {
		t.Error("unable to retrieve images from the database")
	}

	if len(images) == 0 {
		t.Error("expected to get images from the test")
	}
}

func TestUpdateImage(t *testing.T) {
	image.Id = "3"
	imageStore.Save(&image)
	nutImage, err := os.ReadFile("../nut_snickers.png")
	if err != nil {
		t.Error("unable to load image into the db")
	}
	image.Data = nutImage
	if err := imageStore.Update(&image); err != nil {
		t.Errorf("unable to update image %s, err: %s", image.Id, err.Error())
	}
	image.Id = "id does not exist"
	if err := imageStore.Update(&image); err != sql.ErrNoRows {
		t.Error("could not update image with id since it does not exist")
	}

	imageStore.Delete(image.Id)
}

func TestDeleteImage(t *testing.T) {
	image.Id = "3"
	imageStore.Save(&image)

	id, err := imageStore.Delete(image.Id)
	if err != nil {
		t.Errorf("unable to delete image %s, err: %s", image.Id, err.Error())
	}
	if id != image.Id {
		t.Errorf("incorrect id was returned after deleteing an image, got %s, expected %s", id, image.Id)
	}
}
