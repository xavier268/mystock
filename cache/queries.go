package cache

import (
	"fmt"
	"time"

	"github.com/xavier268/mystock/quandl"
)

// getLastRecord retrieve the last available
// record for the selected ticker and measure.
// Refresh as needed.
func (c *Cache) getLastRecord(ticker string, measure Measure) quandl.Record {
	c.refresh(ticker)
	var r quandl.Record

	err := c.db.Where(&quandl.Record{Serie: ticker, Measure: string(measure)}).Order("date desc").First(&r).Error
	if err != nil {
		fmt.Println(ticker, measure)
		panic(err)
	}
	return r
}

// LastPrice retrieve the last closing price,
// returning the date (day) for this price.
func (c *Cache) LastPrice(ticker string) (value float64, when time.Time) {
	la := c.getLastRecord(ticker, Close)
	return la.Value, la.Date
}

// HighLowPrice return the Hihg and Low prices since selected date.
func (c *Cache) HighLowPrice(ticker string, since time.Time) (high, low float64) {
	// TODO
	var h, l float64
	row := c.db.Model(&quandl.Record{}).
		Select(" MAX ( value ) , MIN ( value ) ").
		Where(" measure = ? OR measure = ? OR measure = ? OR measure = ?", Open, Close, High, Low).
		Row()
	row.Scan(&h, &l)
	fmt.Println("Scan result : ", h, l)
	/* if len(hl) != 2 {
	fmt.Println("Scan result : ", hl)
	panic("Unexpected return size")
	} */

	return h, l
}

// ListTickers will list the tickers managed in the cache in a thread safe way.
func (c *Cache) ListTickers() []string {

	c.refGuard.RLock()

	res := make([]string, 0, len(c.ref))
	for s := range c.ref {
		res = append(res, s)
	}

	c.refGuard.RUnlock()

	return res
}
