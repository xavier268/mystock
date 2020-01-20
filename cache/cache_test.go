package cache

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/xavier268/mystock/configuration"
)

// test database file
var ftest = "test.db"

func TestNewMemoryCache(t *testing.T) {
	c := NewMemoryCache(configuration.Load())
	defer c.Close()
}

func TestNewFileCache(t *testing.T) {
	os.Remove(ftest)
	defer os.Remove(ftest)

	c := newCache(configuration.Load(), ftest)
	defer c.Close()

	c.Dump()
	if c.Size() != 0 {
		t.Fail()
	}
}

func TestConstructLocalDB(t *testing.T) {
	// Skip not to leave a db locally.
	// t.Skip()
	fmt.Println("Warning : this test may take a long time to run if the bd is not initialized yet")
	c := NewCache(configuration.Load())
	defer c.Close()

	// Refresh all
	c.Refresh()
}

func TestMostRecent(t *testing.T) {
	c := NewCache(configuration.Load())
	defer c.Close()

	//c.DB.LogMode(true)
	t1 := c.MostRecent("AIR")
	if (t1 == time.Time{}) {
		c.Dump()
		fmt.Println("Most recent AIR time : ", t1)
		t.Fatal("Inconsistent time for AIR")
	}

	t2 := c.MostRecent("NOTHING")
	if (t2 != time.Time{}) {
		t.Fatal("Nil time expected !")
	}

}
