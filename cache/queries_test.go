package cache

import (
	"fmt"
	"testing"
	"time"
)

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
		t.Fatal("unexpected ticker list != 1")
	}

	p, lt := c.LastPrice("ML")
	fmt.Println("Last price : ", p, lt)
}

func TestHighLowPrices(t *testing.T) {
	c := NewCache()
	defer c.Close()

	// Since 5 days
	h1, l1 := c.HighLowPrice("AIR", time.Now().Add(-5*24*time.Hour))
	fmt.Println("HighLow 5 days : ", h1, l1)
	if h1 < l1 {
		fmt.Println("High Low values 5 days: ", h1, l1)
		t.Fatal("Inconsistent highLow")
	}

	// since ever

	he, le := c.HighLowPrice("AIR", time.Time{})
	fmt.Println("HighLow ever : ", he, le)
	if he < le {
		fmt.Println("High Low values ever : ", he, le)
		t.Fatal("Inconsistent highLow")
	}
	if he < h1 || le > l1 {
		fmt.Println("Ever High-Low : ", he, le, " 5 days high-low : ", h1, l1)
		t.Fatal("Inconstitent results. MinMax ever conflicts with MinMax 10 days.")
	}

	// Check last price is consitent
	lp, _ := c.LastPrice("AIR")
	if lp < l1 || lp > h1 {
		t.Fatal("Last price is inconsitent with MinMax")
	}

}
