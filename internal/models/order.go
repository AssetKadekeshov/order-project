package models

type Order struct {
	Id           int    `json:"id"`
	CustomerName string `json:"customerName"`
	Product      string `json:"product"`
	Quantity     int    `json:"quantity"`
	Status       string `json:"status"`
}
