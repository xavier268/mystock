package quandl

import (
	"fmt"
	"time"
)

// Record elementary data structure,
// as retrieved from a source/serie.
type Record struct {
	Date    time.Time `gorm:"primary_key"`
	Source  string
	Serie   string `gorm:"primary_key"`
	Measure string `gorm:"primary_key"`

	Value float64 `gorm:"NOT NULL"`
}

func (r *Record) String() string {
	return fmt.Sprintf("Record : %s/%s %s-%s\t%f",
		r.Source, r.Serie, r.Date.Format(layout),
		r.Measure, r.Value)
}

// RecordProcessor knows how to process a Record
type RecordProcessor func(Record)
