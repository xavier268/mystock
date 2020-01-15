package quandl

import (
	"fmt"
	"time"
)

// Record elementary data structure,
// as retrieved from a source/serie.
// This is the format used to communicate to the external world.
type Record struct {
	Date    time.Time `gorm:"primary_key"`
	Source  string
	Serie   string `gorm:"primary_key"`
	Measure string `gorm:"primary_key"`

	Value float64 `gorm:"NOT NULL"`
}

// Layout used to format time to/from strings.
const Layout string = "2006-01-02"

func (r *Record) String() string {

	return fmt.Sprintf("Record : %s/%s %s-%s\t%f",
		r.Source, r.Serie, r.Date.Format(Layout),
		r.Measure, r.Value)
}

// RecordProcessor knows how to process a Record
type RecordProcessor func(Record)
