package models

type Payment struct {
	OrderTotal string `json:"orderTotal"`
	Products   []struct {
		ProductId    string `json:"productId"`
		Quantity     int    `json:"quantity"`
		Price        int    `json:"price"`
		BoughtInBulk bool   `json:"boughtInBulk"`
	} `json:"products"`
	CustomerId int    `json:"customerId"`
	ReaderId   string `json:"readerId"`
}
