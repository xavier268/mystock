package monitor

// Line defines a line of the shares portfolio
type Line struct {
	// Informative, human readable name
	Name string
	// Unique ticker identifying the  stock shares
	Ticker string
	// Date of purchase - use special type to allow easier unmarshalling
	Date MyTime
	// Historical/average purchase price
	Price float64
	// number of shares
	Volume float64
}

// A Portfolio contains the lines.
type Portfolio struct {
	Lines []Line
}
