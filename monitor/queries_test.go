package monitor

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var m *Monitor

func TestMain(t *testing.M) {
	m = NewMonitor(nil)
	// DEBUG : Setting logmode will display actual sql
	// m.cache.DB.LogMode(true)
	e := t.Run()
	m.Close()
	os.Exit(e)
}

func TestMonitorConstructed(t *testing.T) {
	m.cache.Dump()
}

func TestCountince(t *testing.T) {
	// fmt.Println("Monitor : ", m)
	c1 := m.NewQuery("").Count()
	if c1 <= 100 {
		t.Fatal("Count claims less than 100 records ?")
	}
	c2 := m.NewQuery("AIR").Count()
	if c2 >= c1 {
		t.Fatal("There seems to be only AIR records ?")
	}
	c3 := m.NewQuery("AIR").Since(10 * 24 * time.Hour).Count()
	if c3 >= c2 {
		t.Fatal("Inconsistent since filter")
	}
	fmt.Println("Records in cache : ", c1, c2, c3)
}

func TestDump(t *testing.T) {
	fmt.Println("\nDumping last 5 days of AIR, default order")
	m.NewQuery("AIR").Since(5 * 24 * time.Hour).DumpRecords()
	fmt.Println("\nDumping last 5 days of AIR, desc order")
	m.NewQuery("AIR").OrderDate("desc").Since(5 * 24 * time.Hour).DumpRecords()
	//m.cache.DB.LogMode(true)
	fmt.Println("\nDumping last 5 days of AIR, default order, only TURNOVER")
	m.NewQuery("AIR").Since(5 * 24 * time.Hour).Measure("turNOVEr").DumpRecords()
	fmt.Println("\nDumping last 5 days of AIR, all prices")
	m.NewQuery("AIR").OrderDate("desc").Measure("Open", "Last", "high", "low").Since(5 * 24 * time.Hour).DumpRecords()

}

func TestMinMax(t *testing.T) {
	//m.cache.DB.LogMode(true)
	fmt.Println("\nMinmax  last 5 days of AIR")
	min, max := m.NewQuery("AIR").Since(5 * 24 * time.Hour).MinMax()
	fmt.Println(min, max)
	if min == 0 || max == 0 {
		t.Fatal("Suspicious min max values")
	}
}
