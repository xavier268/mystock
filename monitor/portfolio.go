package monitor

import (
	"fmt"

	"github.com/xavier268/mystock/quandl"
)

// Portfolio is a data structure containing
// current portfolio lines for all tickers.
type Portfolio map[string]VLine

// VLine is a turnover oriented Line.
type VLine struct {
	LastTurnover float64
	HistTurnover float64
	Volume       float64
}

// LoadPortfolio constructs and updates a portfolio object
// from the current Monitor object.
func (m *Monitor) LoadPortfolio() Portfolio {

	var pf Portfolio = make(Portfolio)

	// DEBUG
	// fmt.Println("DEBUG LoadPortfolio, m.lines = ", m.lines)

	for _, l := range m.lines {
		// Load existing earlier values,
		// if multiple lines with same ticker
		var vl VLine = pf[l.Ticker]

		vl.Volume += l.Volume
		vl.HistTurnover += l.Volume * l.Price
		p := m.LastClosingPrice(l.Ticker)
		vl.LastTurnover += l.Volume * p

		// Update portfolio with incremented values.
		pf[l.Ticker] = vl
	}
	return pf
}

// Dump portfolio
func (pf Portfolio) Dump() {
	fmt.Println("Portfolio :")
	for k, v := range pf {
		if v.Volume == 0. {
			fmt.Printf("\t%s  -- no volume --\n", k)
		} else {
			fmt.Printf("\t%s %.2f x %.1f = %.2f -> %.2f x %.1f = %.2f\n",
				k,
				v.HistTurnover/v.Volume,
				v.Volume,
				v.HistTurnover,
				v.LastTurnover/v.Volume,
				v.Volume,
				v.LastTurnover)
		}
	}
	fmt.Println()
}

// Values calculates the total historical (purchase) and last
// values of the portfolio.
func (pf Portfolio) Values() (hist, last float64) {
	for _, v := range pf {
		hist += v.HistTurnover
		last += v.LastTurnover
	}
	return hist, last
}

// LastClosingPrice last known closing price (in cache) for ticker.
func (m *Monitor) LastClosingPrice(ticker string) float64 {
	var rec quandl.Record
	err := m.NewQuery(ticker).Measure("Last").Last(&rec).Error

	if err != nil {
		panic(err)
	}
	return rec.Value
}
