package models

type Product struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	UnitPrice     int64  `json:"unitPrice"`
	BulkPrice     int64  `json:"bulkPrice"`
	Inventory     int    `json:"inventory"`
	ItemsInPacket int    `json:"itemsInPacket"`
	StoreId       string    `json:"storeId"`
}
