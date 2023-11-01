package models

type Payment struct {
	OrderTotal int64         `json:"orderTotal"`
	Products   []ProductOrder `json:"products"`
	CustomerId string         `json:"customerId"`
	ReaderId   string         `json:"readerId"`
}

type ProductOrder struct {
	ProductId    string `json:"productId"`
	Quantity     int    `json:"quantity"`
	Price        int64  `json:"price"`
	BoughtInBulk bool   `json:"boughtInBulk"`
}
