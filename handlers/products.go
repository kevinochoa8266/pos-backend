package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinochoa8266/pos-backend/models"
	"github.com/kevinochoa8266/pos-backend/store"
)

var logger = slog.Default()

var productStore *store.ProductStore

func SetDatabase(db *sql.DB) {
	productStore = store.NewProductStore(db)
	imageStore = store.NewImageStore(db)
	shopStore = store.NewShopStore(db)
	readerStore = store.NewReaderStore(db)
	orderStore = store.NewOrderStore(db)
	customerStore = store.NewCustomerStore(db)
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
		logger.Error(fmt.Sprintf("unable to get product with id %s, err: %s", productId, err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(product)
}

func HandleGetFavorites(w http.ResponseWriter, r *http.Request) {
	favorites, err := productStore.GetFavorites()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("unable to send the favorite products to the client", "err:", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(favorites)
}

func HandleAddProduct(writer http.ResponseWriter, request *http.Request) {
	var product models.Product
	json.NewDecoder(request.Body).Decode(&product)

	id, err := productStore.Save(&product)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		logger.Error(err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(id)
}

func HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	_, err := productStore.Get(product.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := productStore.Update(&product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(err.Error())
		return
	}

	if err := productStore.Delete(&product); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
