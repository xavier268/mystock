package qdl

import (
	"time"
)

// Record elementary data structure.
// Used to store in database cache.
type Record struct {
	Date    time.Time `gorm:"primary_key"`
	Ticker  string    `gorm:"primary_key"`
	Measure string    `gorm:"primary_key"`

	Value float64 `gorm:"NOT NULL"`
}

// RecordProcessor knows how to process a Record
type RecordProcessor func(*Record)
