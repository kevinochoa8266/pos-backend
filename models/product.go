package models

type Product struct {
	Id            string
	Name          string
	Price         int
	BulkPrice     int
	Inventory     int
	ItemsInPacket int
	StoreId       int
}
