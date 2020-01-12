package qdl

import (
	"strings"
)

// Code defines the quandl-database and the time series (ticker).
type Code struct {
	DB     string // ex : EURONEXT
	Ticker string // ex : AYDEN
}

func (c *Code) String() string {
	return strings.ToUpper(c.DB + "/" + c.Ticker)
}
