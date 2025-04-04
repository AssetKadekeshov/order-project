package models

type OrderEdit struct {
	CustomerName string `json:"customerName"`
	Product      string `json:"product"`
	Quantity     int    `json:"quantity"`
	Status       string `json:"status"`
}
