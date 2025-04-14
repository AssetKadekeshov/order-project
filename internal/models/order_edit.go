package models

import "github.com/lib/pq"

type OrderEdit struct {
	CustomerName string         `json:"customerName"`
	Products     pq.StringArray `json:"products"`
	Status       string         `json:"status"`
}
