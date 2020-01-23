package monitor

import (
	"errors"
	"fmt"
)

// CheckGainLoss provides a check for total value percent change.
// Percent is expressed as a decimal number,
// ie 0.3 means 30% change threshold, not 0.3% !
func CheckGainLoss(percent float64) Check {
	return func(m *Monitor) (mess string, e error) {
		p := m.LoadPortfolio()
		hist, last := p.Values()
		if hist <= 0 || last <= 0 {
			fmt.Printf("Error in CheckGainLoss : unexpected porfolio values : min = %v, max = %v\n", hist, last)
			return "", errors.New("checkGainLoss generated unexpected porfolio values")
		}
		if last >= (1+percent)*hist || last <= (1-percent)*hist {
			mess = fmt.Sprintf("Portfolio change exceeded threshold : %.2f -> %.2f (%.2f%%)", hist, last, 100*(last-hist)/hist)
			return mess, nil
		}
		return "", nil
	}
}

// CheckPriceChange alerts if price change more than threshold.
// Percent is expressed as a decimal number,
// ie 0.3 means 30% change threshold, not 0.3% !
func CheckPriceChange(ticker string, percent float64) Check {
	return func(m *Monitor) (mess string, e error) {
		p := m.LoadPortfolio()

		vl := p[ticker]
		hist, last := vl.HistTurnover, vl.LastTurnover
		if last >= (1+percent)*hist || last <= (1-percent)*hist {
			mess = fmt.Sprintf("Price of <%s> change exceeded threshold : %.2f -> %.2f (%.2f%%)", ticker, hist, last, 100*(last-hist)/hist)
			return mess, nil
		}
		return "", nil
	}
}

// CheckPriceChangeAll will check for price change line by line.
// retun a function that will aggregate the per-ticker functions.
func CheckPriceChangeAll(percent float64) Check {
	return func(m *Monitor) (mess string, e error) {
		p := m.LoadPortfolio()
		for t := range p {
			s, e := CheckPriceChange(t, percent)(m)
			if e != nil {
				return mess, e
			}
			mess = "\n" + s
		}
		return mess, nil
	}
}
