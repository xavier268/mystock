package quandl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Q is the QUANDL query structure.
// It is thread safe.
type Q struct {
	version string
	query   url.Values
	source  Source
}

func (q *Q) Version() string {
	return q.version
}

// New construct a new Q with the specified options.
// The source (eg : EURONEXT) is regquired.
// Existing Q may be used multiple times
// to query different time series from the same Source.
func New(source Source, options ...QOption) *Q {

	if len(source) == 0 {
		panic("quandl source cannot be empty")
	}
	q := new(Q)
	q.source = Source(strings.ToUpper(string(source)))

	q.version = "https://www.quandl.com/api/v3/"

	// apply the options ...
	for _, opt := range options {
		opt(q)
	}

	// add the key
	q.query = make(url.Values)
	q.query.Set("api_key", APISecretKey())
	return q
}

// Source represent a source database in QUANDL.
type Source string

// ApiSecretKey is a hook function that provides
// the QUANDL api key upon request.
// Redefine this function to provide your APISecretKey.
// For instance, as an init :
//
//  func init() {
// 	fmt.Println("Initialising API secret key")
// 	APISecretKey = func() string {
// 		return "mysecretkey"
// 	}
// }
var APISecretKey func() string

// layout used to format time to/from strings.
const layout string = "2006-01-02"

// WalkDataset will collect Records from
// the source/ticker and process each
// record with the provided processor.
func (q *Q) WalkDataset(
	ticker string,
	processor RecordProcessor) {

	// If processor is nil,
	// just dump records.
	if processor == nil {
		processor = func(r Record) {
			fmt.Println(r)
		}
	}

	// make url string
	u := q.version +
		"datasets/" +
		string(q.source) +
		"/" + strings.ToUpper(ticker) +
		".json?" +
		q.query.Encode()

	// debug
	fmt.Println(u)

	// Get response from web service
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected http status :", resp.StatusCode, resp.Status)
		panic(resp.Status)
	}

	// Decode body into datasetModel
	defer resp.Body.Close()
	var d datasetModel
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		panic(err)
	}

	// debug
	// fmt.Println(d)

	// Check expected columns are where
	// they should be ...
	/*
		ec := expectedColumns()
		if len(d.DatasetData.ColumnNames) != len(ec) {
			panic("Unexpected columns format !")
		}
		for _, s := range d.DatasetData.ColumnNames {
			if !ec[s] {
				panic("Unexpected column name : " + s)
			}
		}
	*/
	// Expect first column to be date.
	if len(d.Dataset.ColumnNames) == 0 ||
		d.Dataset.ColumnNames[0] != "Date" {
		panic("Date is expected to be the first column in column names")
	}

	for i := range d.Dataset.Data {
		dd := d.Dataset.Data[i]

		// get record time
		t := dd[0].(string)
		tt, err := time.Parse(layout, t)
		if err != nil {
			panic(err)
		}

		// create/set records
		for c, s := range d.Dataset.ColumnNames {
			// create a record per value, excpet for Date
			if c != 0 { // avoid Date ...
				r := new(Record)
				r.Source = string(q.source)
				r.Serie = ticker
				r.Date = tt
				r.Measure = s
				r.Value = dd[c].(float64)
				processor(*r) // emit or process created record
			}
		}
	}

}
