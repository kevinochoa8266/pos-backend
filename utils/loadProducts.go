package utils

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

func LoadProductsIntoStore(storeId int64, db *sql.DB, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	ps := store.NewProductStore(db)

	for _, line := range data[1:50] {
		product, err := extractProduct(line)
		if err != nil {
			panic(err)
		}
		product.StoreId = int(storeId)
		if _, err = ps.Save(product); err != nil {
			return fmt.Errorf("product (%s, %s) could not be saved: %s",
				product.Id,
				product.Name,
				err.Error())
		}
	}
	return nil
}

func extractProduct(l []string) (*models.Product, error) {
	var product models.Product
	product.Id = l[0]
	product.Name = l[1]
	tax, _ := strconv.Atoi(l[2])
	taxDecimal := float32(tax) / 100.0
	taxApplied, _ := strconv.ParseBool(l[3])
	product.Inventory, _ = strconv.Atoi(l[4])
	product.BulkPrice, _ = strconv.Atoi(l[5])
	product.UnitPrice, _ = strconv.Atoi(l[6])
	product.ItemsInPacket, _ = strconv.Atoi(l[7])

	if tax != 0 && !taxApplied {
		var taxPrice float32 = float32(product.BulkPrice) * taxDecimal
		product.BulkPrice += int(taxPrice)
		if product.UnitPrice != 0 {
			product.UnitPrice += int(taxPrice)
		}
	}

	if product.ItemsInPacket != 0 {
		product.Inventory *= product.ItemsInPacket
	}

	return &product, nil
}