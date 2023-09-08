package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

func main() {
	f, err := os.Open("candy_data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.GetConnection("store.db"); if err != nil {
		log.Fatal(err)
	}
	ps := store.NewProductStore(db)

	for i, line := range data {
		if i > 0 {
			var product models.Product
			var tax int
			var taxDecimal float32
			var taxApplied bool
			var itemsInPacket int


			for j, field := range line {
				switch j {
				case 0:
					product.Id = field
					continue
				case 1:
					product.Name = field
					continue
				case 2:
					tax, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal(err)
					}
					taxDecimal = float32(tax) / 100
					continue
				case 3:
					taxApplied, err := strconv.ParseBool(field)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(taxApplied)
					continue
				case 4:
					inventory, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal(err)
					}
					product.Inventory = inventory
					continue
				case 5:
					bulkPrice, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal(err)
					}
					product.BulkPrice = bulkPrice
					continue
				case 6:
					individualPrice, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal(err)
					}
					product.Price = individualPrice
					continue
				case 7:
					itemsInPacket, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal(err)
					}
					product.ItemsInPacket = itemsInPacket
				}
				if tax != 0 && !taxApplied {
					var taxPrice float32 = float32(product.BulkPrice) * taxDecimal
					product.BulkPrice += int(taxPrice)
					if product.Price != 0 {
						product.Price += int(taxPrice)
					}
				}

				if itemsInPacket != 0 {
					product.Inventory *= itemsInPacket
				}
				
				_, err := ps.Save(&product); if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
