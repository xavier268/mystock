package qdl

import (
	"fmt"
	"testing"
	"time"
)

func TestInsertOrder(t *testing.T) {
	c := NewMemoryDB()
	defer c.Close()

	if r, o := c.Count(); r != 0 && o != 0 {
		t.Fail()
	}
	o := new(Order)
	o.Ticker = "ML"
	o.Date = time.Now()
	o.Price = 110.3
	o.Volume = 10
	c.CreateOrder(o)
	c.Dump()
	if _, n := c.Count(); n != 1 {
		fmt.Println("Order count : ", n)
		panic("Wrong db count !")
	}

}
