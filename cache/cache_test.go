package cache

import (
	"fmt"
	"os"
	"testing"
)

// test database file
var ftest = "test.db"

func TestNewMemoryCache(t *testing.T) {
	c := NewMemoryCache(mysecret)
	defer c.Close()
}

func TestNewFileCache(t *testing.T) {
	os.Remove(ftest)
	c := newCache(mysecret, ftest)
	defer c.Close()
	c.Dump()
	if c.Size() != 0 {
		t.Fail()
	}
	c.Close() // before removing file ...
	os.Remove(ftest)
}

func TestConstructLocalDB(t *testing.T) {
	// Skip not to leave a db locally.
	// t.Skip()
	fmt.Println("Warning : this test may take a long time to run if the bd is not initialized yet")
	c := NewCache(mysecret)
	defer c.Close()

	// Fill wil ML & AIR
	c.LastPrice("AIR")
	c.LastPrice("ML")
}
