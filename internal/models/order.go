package models

import (
	"github.com/lib/pq"
)

type Order struct {
	Id            int            `json:"id"`
	CustomerName  string         `json:"customerName"`
	Products      pq.StringArray `json:"products" gorm:"type:text[]"`
	TotalQuantity int            `json:"totalQuantity"`
	Status        string         `json:"status"`
}
