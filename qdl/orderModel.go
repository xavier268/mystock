package qdl

import (
	"time"
)

// Order is a portfolio change line (sell or buy)
type Order struct {
	ID     uint `gorm:"primary_key"`
	Ticker string
	Date   time.Time // Purchase date
	Price  float64   // Purchase price
	Volume float64   // positive to buy, negative to sell
}
