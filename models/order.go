package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id           uuid.UUID `json:"id"`
	ProductId    string    `json:"productId"`
	CustomerId   int       `json:"customerId"`
	Date         time.Time `json:"date"`
	BoughtInBulk bool      `json:"boughtInBulk"`
	Quantity     int       `json:"quantity"`
	TotalPrice   int       `json:"totalPrice"`
}
