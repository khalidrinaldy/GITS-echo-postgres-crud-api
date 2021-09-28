package models

type Movie struct {
	Id     int `json:"id"`
	Title  string `json:"title"`
	Rating int `json:"rating"`
}

type Movies struct {
	Movies []Movie	`json:"movies"`
}