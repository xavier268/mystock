package qdl

import "net/url"

// QDL is the main object to query Quandl web services.
type QDL struct {
	url   string
	query url.Values
}

// NewQDL  create a new QDL
func NewQDL() *QDL {
	var e error
	q := new(QDL)
	q.url = "https://www.quandl.com/api/v3/datasets/"
	if e != nil {
		panic(e)
	}

	q.query = make(url.Values)
	q.query.Set("api_key", APISecretKey)
	return q
}
