package models

import "time"

type Payment struct {
	OrderId    string `json:"id"`
	OrderTotal string `json:"orderTotal"`
	Products   []struct {
		ProductId    string `json:"productId"`
		Quantity     int    `json:"quantity"`
		Price        int    `json:"price"`
		BoughtInBulk bool   `json:"boughtInBulk"`
	} `json:"products"`
	CustomerId int       `json:"customerId"`
	Date       time.Time `json:"date"`
	ReaderId   string    `json:"readerId"`
}