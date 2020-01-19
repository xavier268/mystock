package monitor

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/xavier268/mystock/quandl"
)

// Query is a warpper around gorm.DB
type Query struct {
	*gorm.DB
}

// NewQuery starts to build a new record query
// for the provided ticker. If provided ticker
// is an empty string, cover the entire base.
func (m *Monitor) NewQuery(ticker string) *Query {
	q := m.cache.DB.Model(&quandl.Record{})
	if len(ticker) != 0 {
		q = q.Where(" serie = ? ", strings.ToUpper(ticker))
	}
	return &Query{q}
}

// Since limit the search from that duration in the past.
// EX : Since( 30 * 24 * time.Hour) to search past 30 days.
func (q *Query) Since(d time.Duration) *Query {
	t := time.Now().Add(-d)
	qq := q.Where("date >= ?", t)
	return &Query{qq}
}

// Order specify the sort order.
type Order string

// Measure specify how to order
type Measure string

// OrderDate specify order.
// Valid Order are : ASC or DESC
// It is case insensitive.
// Invalid values are silently ignored.
func (q *Query) OrderDate(order Order) *Query {
	od := strings.ToUpper(string(order))
	if od != "ASC" &&
		od != "DESC" {
		return q
	}
	return &Query{q.DB.Order(" date " + od)}
}

// Limit results
func (q *Query) Limit(n int) *Query {
	return &Query{q.DB.Limit(n)}
}

// Select manually select fields
func (q *Query) Select(fields ...string) *Query {
	if len(fields) == 0 {
		return q
	}
	return &Query{q.DB.Select(fields)}
}

// Count the number of records selected.
func (q *Query) Count() float64 {
	var c float64
	err := q.DB.Count(&c).Error
	if err != nil {
		panic(err)
	}
	return c
}

// DumpRecords the selected rows.
// Assume record format.
func (q *Query) DumpRecords() {
	var rec quandl.Record

	rw, err := q.DB.Rows()
	defer rw.Close() // Important !!
	if err != nil {
		panic(err)
	}
	for rw.Next() {
		q.DB.ScanRows(rw, &rec)
		fmt.Println(rec.String())
	}
}

// Dump the selected rows.
// Makes no format assumption
func (q *Query) Dump() {
	var rec interface{}

	rw, err := q.DB.Rows()
	defer rw.Close() // Important !!
	if err != nil {
		panic(err)
	}
	for rw.Next() {
		q.DB.ScanRows(rw, &rec)
		fmt.Println(rec)
	}
}

// MinMax values for selected query.
func (q *Query) MinMax() (min, max float64) {
	err := q.DB.Select(" MIN(value)  ,  MAX(value) ").Row().Scan(&min, &max)
	if err != nil {
		panic(err)
	}
	fmt.Println(min, max)
	return min, max
}

// Measure selects  any of the provided  measures.
// (Measure are ORed).
func (q *Query) Measure(ms ...string) *Query {

	if len(ms) == 0 {
		return q
	}
	var mm []string
	for _, s := range ms {
		mm = append(mm, strings.ToUpper(s))
	}
	return &Query{q.DB.Where("Measure IN  ( ? )", mm)}
}
