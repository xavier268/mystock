package cache

// Measure defines the various values taht can be retrieved from the database.
type Measure string

// Predefined measures
const (
	Open     Measure = "OPEN"
	Close    Measure = "LAST" // alias
	Last     Measure = "LAST"
	High     Measure = "HIGH"
	Low      Measure = "LOW"
	Volume   Measure = "VOLUME"
	Turnover Measure = "TURNOVER"
)
