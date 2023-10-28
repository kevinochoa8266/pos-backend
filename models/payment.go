package models

type Payment struct {
	OrderTotal string `json:"orderTotal"`
	Products   []struct {
		ProductId    string `json:"productId"`
		Quantity     int    `json:"quantity"`
		Price        int64    `json:"price"`
		BoughtInBulk bool   `json:"boughtInBulk"`
	} `json:"products"`
	CustomerId string    `json:"customerId"`
	ReaderId   string `json:"readerId"`
}
