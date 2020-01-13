package qdl

import (
	"fmt"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	c := NewMemoryDB()
	defer c.Close()

	if r, o := c.Count(); r != 0 && o != 0 {
		t.Fail()
	}
	o := new(Order)
	o.Ticker = "ML"
	o.Date = time.Now()
	o.Price = 100
	o.Volume = 10
	c.CreateOrder(o)
	c.Dump()
	if _, n := c.Count(); n != 1 {
		fmt.Println("Order count : ", n)
		panic("Wrong db count !")
	}
	o = new(Order)
	o.Ticker = "ML"
	o.Date = time.Now()
	o.Price = 50
	o.Volume = 5
	c.CreateOrder(o)
	if _, n := c.Count(); n != 2 {
		fmt.Println("Order count : ", n)
		panic("Wrong db count !")
	}
	r := c.PortfolioHistValue()
	if r != 5*50+10*100 {
		panic("Wrong portfolio value")
	}
}
