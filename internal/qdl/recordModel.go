package qdl

import "time"

// Record elementary data structure
type Record struct {
	Date    time.Time
	Ticker  string
	Measure string
	Value   float64
}

// RecordProcessor knows how to process a Record
type RecordProcessor func(Record)
