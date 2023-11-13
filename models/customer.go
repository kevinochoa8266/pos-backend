package models

type Customer struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}
