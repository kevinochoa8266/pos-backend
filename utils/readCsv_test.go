package utils_test

import (
	"testing"

	"github.com/kevinochoa8266/pos-backend/utils"
)

func TestReadCsv(t *testing.T) {
	if err := utils.ReadCsvData("../candy_data.csv", ":memory:"); err != nil {
		t.Errorf("unable to save the products into the test db. %s", err.Error())
	}

	if err := utils.ReadCsvData("random Path", ":memory:"); err == nil {
		t.Error("there should be no file at the given path to read")
	}
}
