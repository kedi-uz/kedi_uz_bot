package models


type AnimalType struct {
	Id int `json:"id"`
	Title string `json:"title"`
}

type GenderType struct {
	Id int `json:"id"`
	Title string `json:"title"`
}

type PostedBy struct {
	Id int `json:"id"`
	UserName string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type District struct {
	Id int `json:"id"`
	Title string `json:"title"`
}

type LostAnimal struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Photo string `json:"photo"`
	Description string `json:"description"`
	Location string `json:"location"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	DateLost string	`json:"date_lost"`
	PhoneNumber string	`json:"phone_number"`
	AnimalType AnimalType `json:"anima_type"`
	GenderType GenderType	`json:"gender_type"`
	PostedBy PostedBy	`json:"posted_by"`
	District District	`json:"district"`
}