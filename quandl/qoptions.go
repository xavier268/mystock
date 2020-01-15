package quandl

import (
	"fmt"
	"time"
)

// A QOption is an option that can be
// provided when constructing a new Q.
type QOption func(*Q)

// OptionStartDate sets the start date.
func OptionStartDate(when time.Time) QOption {
	return func(q *Q) {
		q.query.Set("start_date", when.Format(layout))
	}
}

// OptionEndDate sets the end date.
func OptionEndDate(when time.Time) QOption {
	return func(q *Q) {
		q.query.Set("end_date", when.Format(layout))
	}
}

// OptionOrderAsc ascending order.
var OptionOrderAsc QOption = func(q *Q) {
	q.query.Set("order", "asc")
}

// OptionOrderDesc descending order.
var OptionOrderDesc QOption = func(q *Q) {
	q.query.Set("order", "desc")
}

// OptionFrequency sets the sampling frequency.
func OptionFrequency(freq Frequency) QOption {
	return func(q *Q) {
		q.query.Set("collapse", string(freq))
	}
}

// Frequency of the samples.
type Frequency string

// Frequency constants.
const (
	Daily     Frequency = "daily"
	Weekly    Frequency = "weekly"
	Monthly   Frequency = "montly"
	Quarterly Frequency = "quarterly"
	Annual    Frequency = "annual"
)

// A Transform defines transformations
// to apply to data.
type Transform string

// Transform constants.
const (
	Delta         Transform = "diff"
	DeltaPercent  Transform = "rdiff"
	CumulativeSum Transform = "cumul"
	Base100       Transform = "normalize"
)

// OptionLimit sets the number of dataset returned.
// Set limit to 1 to return the last dataset.
func OptionLimit(n int) QOption {
	return func(q *Q) {
		q.query.Set("limit", fmt.Sprintf("%d", n))
	}
}
