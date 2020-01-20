package quandl

// datasetModel used internally to parse the data returned by GDL
type datasetModel struct {
	Dataset struct {
		Limit       *string         `json:"limit"`
		Transform   *string         `json:"transform"`
		ColumnIndex *int            `json:"column_index"`
		ColumnNames []string        `json:"column_names"`
		StartDate   string          `json:"start_date"`
		EndDate     string          `json:"end_date"`
		Frequency   string          `json:"frequency"`
		Data        [][]interface{} `json:"data"`
	} `json:"dataset"`
	QuandlError struct {
		Code    string `jsn:"code"`
		Message string `jsn:"message"`
	} `json:"quandl_error"`
}

// expectedColumns provides a map of the columns expected to be found.
func expectedColumns() map[string]bool {
	m := make(map[string]bool)
	for _, k := range [...]string{
		"Date",
		"Open",
		"Last",
		"High",
		"Low",
		"Volume",
		"Turnover",
	} {
		m[k] = true
	}
	return m
}
