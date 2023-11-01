package models

import (
	"time"
)

type Order struct {
	Id                     string    `json:"id"`
	ProductId              string    `json:"productId"`
	CustomerId             string    `json:"customerId"`
	Date                   time.Time `json:"date"`
	BoughtInBulk           bool      `json:"boughtInBulk"`
	Quantity               int       `json:"quantity"`
	ProductPriceAtPurchase int64     `json:"price"`
}

// delete BoughtInBulk from the order model
