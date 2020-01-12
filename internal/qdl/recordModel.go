package qdl

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Record elementary data structure.
type Record struct {
	gorm.Model
	Date    time.Time `gorm:"INDEX;NOT NULL"`
	Ticker  string    `gorm:"INDEX;NOT NULL"`
	Measure string    `gorm:"INDEX;NOT NULL"`
	Value   float64   `gorm:"NOT NULL"`
}

// RecordProcessor knows how to process a Record
type RecordProcessor func(Record)
