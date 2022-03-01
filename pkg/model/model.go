package model

import "time"

type Email string
type Phone string

type Contact struct {
	Email       Email `json:"email"`
	PhoneNumber Phone `json:"phone"`
}

type Address struct {
	ZipCode string `json:"zipcode"`
	Street  string `json:"street"`
	Country string `json:"country"`
	City    string `json:"city"`
}

type Customer struct {
	Id        int64       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	BirthDate time.Time `json:"birthdate"`
	Address   Address
	Contact   Contact
}
