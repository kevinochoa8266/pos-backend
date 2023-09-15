package main

import (
	"os"

	"github.com/kevinochoa8266/pos-backend/utils"
)

func main() {

	if _, err := os.Open("store.db"); err != nil {
		if err := utils.ReadCsvData("candy_data.csv", "store.db"); err != nil {
			panic(err)
		}
	}

	

}
