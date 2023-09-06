package models

type Product struct {
	Id            string
	Name          string
	Price         float32
	BulkPrice     float32
	Inventory     int
	BulkInventory int
	HasTax        bool
	TaxIncluded   bool
	TaxRate       float32
	StoreId       int
}
