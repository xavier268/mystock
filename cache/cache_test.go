package cache

import (
	"fmt"
	"os"
	"testing"

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
	c := newCache(configuration.Load(), ftest)
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
	c := NewCache(configuration.Load())
	defer c.Close()

	// Refresh all
	c.Refresh()
}
