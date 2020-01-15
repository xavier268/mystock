package quandl

import "time"

// QOtions are options that can be
// defined when constructing a new Q.
type QOption func(*Q)

func OptionStartDate(when time.Time) {
	// TODO
}

func OptionEndDate(when time.Time) {
	// TODO
}
