package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinochoa8266/pos-backend/store"

	"log/slog"
)

var db *sql.DB
var productStore *store.ProductStore

var logger = slog.Default()

func init() {
	dbUrl := "store.db"

	var err error

	db, err = store.GetConnection(dbUrl)
	if err != nil {
		panic(err)
	}
	productStore = store.NewProductStore(db)
}

func HandleGetProducts(writer http.ResponseWriter, request *http.Request) {

	products, err := productStore.GetAll()
	if err != nil {
		logger.Error("could not retrieve products from the database. %s", err.Error(), 1)
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(products)
}

func HandleGetProduct(writer http.ResponseWriter, request *http.Request) {
	productId := mux.Vars(request)["id"]

	product, err := productStore.Get(productId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		logger.Error("could not retrieve product with id %s due to: %s", productId, err.Error())
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(product)
}
