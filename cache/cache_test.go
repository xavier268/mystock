package cache

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// test database file
var ftest = "test.db"

func TestNewMemoryCache(t *testing.T) {
	c := NewMemoryCache()
	defer c.Close()
}

func TestNewFileCache(t *testing.T) {
	os.Remove(ftest)
	c := newCache(ftest)
	defer c.Close()
	c.Dump()
	if c.Size() != 0 {
		t.Fail()
	}

}

func TestRetrieveValue(t *testing.T) {
	c := NewMemoryCache()
	defer c.Close()

	tt := c.ListTickers()
	if len(tt) != 0 {
		fmt.Println(tt)
		t.Fatal("unexcpected ticker list != 0")
	}

	r := c.getLastRecord("ML", Last)
	if r.Serie != "ML" || r.Measure != string(Last) {
		t.Fatal("Record returned does not answer request !")
	}
	fmt.Println(r.String())
	c.Dump()
	if c.Size() < 1000 {
		t.Fatal("Database should have at least 1000 records !")
	}
	if r.Date.Before(time.Now().Add(-48 * time.Hour)) {
		t.Fatal("The value retuned is older than 48 hours")
	}
	if ti, ok := c.ref["ML"]; !ok ||
		ti.Add(30*time.Second).Before(time.Now()) ||
		len(c.ref) != 1 {
		fmt.Println("Refresh map : ", c.ref)
		t.Fatal("Last refresh should have happenend less than 30 sec ago")
	}

	tt = c.ListTickers()
	if len(tt) != 1 {
		fmt.Println("Ticker list : ", tt, "size : ", len(tt))
		t.Fatal("unexcpected ticker list != 1")
	}
}