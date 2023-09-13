package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

func ReadCsvData(path string) error { //TODO: add better error handling and finish this up tomorrow.
	f, err := os.Open("candy_data.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.GetConnection("store.db")

	if err != nil {
		log.Fatal(err)
	}
	store.CreateSchema(db)

	query := "INSERT INTO STORE (id, name, address) VALUES(?,?,?)"
	_, err = db.Exec(query, 1, "casa dulce", "274 Cape Harbour")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ps := store.NewProductStore(db)

	for _, line := range data[1:5] {
		product, err := extractProduct(line)
		if err != nil {
			panic(err)
		}
		product.StoreId = 1
		if _, err = ps.Save(product); err != nil {
			panic(err)
		}
	}
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
	product.Price, _ = strconv.Atoi(l[6])
	product.ItemsInPacket, _ = strconv.Atoi(l[7])

	if tax != 0 && !taxApplied {
		var taxPrice float32 = float32(product.BulkPrice) * taxDecimal
		product.BulkPrice += int(taxPrice)
		if product.Price != 0 {
			product.Price += int(taxPrice)
		}
	}

	if product.ItemsInPacket != 0 {
		product.Inventory *= product.ItemsInPacket
	}

	return &product, nil
}
