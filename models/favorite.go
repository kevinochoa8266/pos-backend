package models

type Favorite struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	UnitPrice int64  `json:"unitPrice"`
	BulkPrice int64  `json:"bulkPrice"`
	Image     []byte `json:"image"`
}
