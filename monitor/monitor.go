package monitor

import (
	"context"

	"github.com/xavier268/mystock/cache"
	"github.com/xavier268/mystock/configuration"
)

// Monitor allows to monitor share price evolution and send alerts.
type Monitor struct {
	// the portfolio we monitor
	porfolio []configuration.Line
	// how we access market and historical data and prices
	cache cache.Cache
	// The checks we have to perform while monitoring.
	checks []Check
	// How do we send messages to alert us about stock price changes
	// Should ignore empty messages.
	alert Alert
	// APIScretKey for quandl
	secretQuandl string
	// SNS Alert Notification topic
	topicSNS string
}

// Check performs regular checks on the portfolio.
// If the current situation needs to alert the user, return a message.
// or an empty string if no message needed, no attention needed.
// Return non nil error if unable to conduct the check.
type Check func(Monitor) (message string, err error)

// Alert is a function that will alert me if special conditions are detected upon checking.
type Alert func(interface{}) error

// NewMonitor creates and initialize a new Monitor object.
// The portfolio is initialized from the configuration file.
// The cache is created, and the local databse created as needed.
func NewMonitor(alert Alert, checks ...Check) *Monitor {
	m := new(Monitor)
	m.topicSNS = "arn:aws:sns:eu-west-1:057259519792:mystock"

	panic("Not implemented")
	// return m
}

// Run the Monitor, performing regular checks.
// Upon context cancel or timeout, the Monitor will be closed.
func (m *Monitor) Run(context context.Context) {
	panic("Not implemented")
}

// Close the Monitor object.
func (m *Monitor) Close() {
	if m.alert != nil {
		m.alert("The monitoring is now stopped.")
	}
	m.cache.Close()
}
