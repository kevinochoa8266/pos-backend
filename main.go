package main

import "github.com/kevinochoa8266/pos-backend/utils"

func main() {

	if err := utils.ReadCsvData("candy_data.csv", "store.db"); err != nil {
		panic(err)
	}

}
