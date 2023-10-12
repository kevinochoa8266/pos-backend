package models

type Image struct {
	Id   string `json:"id"`
	Data []byte `json:"data"`
}
