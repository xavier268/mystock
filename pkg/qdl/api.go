package qdl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// QDL is the main object to query Quandl web services.
type QDL struct {
	url   string
	query url.Values
}

// New create a new QDL.
// Databse and series coeds must be specified.
func New() *QDL {
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

// AddCode add a relative path to the url, keeping the query elements.
// Examples : AddCode("EURONEXT")
func (q *QDL) makeURL(dbcode, scode string) *url.URL {
	u := q.url + strings.ToUpper(dbcode) + "/" + strings.ToUpper(scode) + "/data.json?"
	// Add the existing query data ...
	u += q.query.Encode()
	url, e := url.Parse(u)
	if e != nil {
		panic(e)
	}
	return url
}

// GetHistory retrieve historical time serie for provided codes.
func (q *QDL) GetHistory(dbcode, scode string) *History {

	// launch http request
	u := q.makeURL(dbcode, scode)
	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}

	// check no error ?
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request : %s\nStatus Code error : %s\n", u.String(), resp.Status)
		panic(resp.Status)
	}

	// retrieve json body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// debug
	// fmt.Printf("Returned data : %s\n", body)

	// convert body to History
	h := q.ToHistory(scode, body)

	// debug
	// fmt.Printf("\nHistory : \t%+v\n", *h)

	return h

}
