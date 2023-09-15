package models

type Product struct {
	Id            string
	Name          string
	UnitPrice     int
	BulkPrice     int
	Inventory     int
	ItemsInPacket int
	StoreId       int
}
