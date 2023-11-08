package models

type Reader struct {
	Id         string `json:"readerId"`
	Name       string
	LocationId string
}
