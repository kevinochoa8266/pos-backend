package models

import (
	"time"
)

type Order struct {
	Id              string    `json:"id"`
	Date            time.Time `json:"date"`
	Quantity        int       `json:"quantity"`
	PriceAtPurchase int64     `json:"priceAtPurchase"`
	ProductId       string    `json:"productId"`
	CustomerId      int       `json:"customerId"`
}
