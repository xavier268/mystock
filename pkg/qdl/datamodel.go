package qdl

import (
	"encoding/json"
	"time"
)

// DataModel data returned by GDL
// used internally to produce an History.
type datasetModel struct {
	DatasetData struct {
		Limit       *string         `json:"limit"`
		Transform   *string         `json:"transform"`
		ColumnIndex *int            `json:"column_index"`
		ColumnNames []string        `json:"column_names"`
		StartDate   string          `json:"start_date"`
		EndDate     string          `json:"end_date"`
		Frequency   string          `json:"frequency"`
		Data        [][]interface{} `json:"data"`
	} `json:"dataset_data"`
}

// History is a parsed model of the received data for a given code.
type History struct {
	Code    string
	Columns []string
	Data    []Record
}

// Record elementary data structure
type Record struct {
	Date   time.Time
	Values map[string]float64
}

// ToHistory will decode the provided json body into an History data structure.
func (q *QDL) ToHistory(code string, body []byte) *History {
	const layout string = "2006-01-02"

	d := new(datasetModel)
	err := json.Unmarshal(body, d)
	if err != nil {
		panic(err)
	}

	h := new(History)
	h.Code = code
	h.Columns = d.DatasetData.ColumnNames
	if h.Columns[0] != "Date" {
		panic("Expecting Date as first column !")
	}

	for i := range d.DatasetData.Data {
		dd := d.DatasetData.Data[i]
		var r Record

		// set record time
		t := dd[0].(string)
		r.Date, err = time.Parse(layout, t)
		if err != nil {
			panic(err)
		}

		// set record values
		r.Values = make(map[string]float64)
		for c, s := range h.Columns {
			if c != 0 { // avoid Date
				r.Values[s] = dd[c].(float64)
			}
		}
		h.Data = append(h.Data, r)
	}

	return h
}
