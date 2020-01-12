package qdl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Refresh will refresh the historic data for the provided code
// and apply the RecordProcessor to each Record.
func (q *QDL) Refresh(code Code, process RecordProcessor) {

	const layout string = "2006-01-02"

	u := q.makeURLString(code)

	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}

	// check no error ?
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request : %s\nStatus Code error : %s\n", u, resp.Status)
		panic(resp.Status)
	}

	// retrieve json body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse body into datasetModel
	d := new(datasetModel)
	err = json.Unmarshal(body, d)
	if err != nil {
		panic(err)
	}

	// Check expected columns are where they should be ...
	ec := expectedColumns()
	if len(d.DatasetData.ColumnNames) != len(ec) {
		panic("Unexpected columns format !")
	}
	for _, s := range d.DatasetData.ColumnNames {
		if !ec[s] {
			panic("Unexpected column name : " + s)
		}
	}
	if d.DatasetData.ColumnNames[0] != "Date" {
		panic("Date is expected to be the first column in column names")
	}

	for i := range d.DatasetData.Data {
		dd := d.DatasetData.Data[i]

		// get record time
		t := dd[0].(string)
		tt, err := time.Parse(layout, t)
		if err != nil {
			panic(err)
		}

		// create/set records
		for c, s := range d.DatasetData.ColumnNames {
			// create a record per value, excpet for Date
			if c != 0 { // avoid Date ...
				r := Record{}
				r.Ticker = code.Ticker
				r.Date = tt
				r.Measure = s
				r.Value = dd[c].(float64)
				process(r) // emit or process created record
			}
		}
	}

}

// makeURL produce the webservice URL request for the given code.
func (q *QDL) makeURLString(code Code) string {
	u := q.url + code.String() + "/data.json?"
	// Add the existing query data ...
	u += q.query.Encode()
	return u
}
