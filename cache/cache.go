package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/xavier268/mystock/quandl"

	// blank import used for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Cache is the DB access object used to cach or retrieve data.
type Cache struct {
	// database for caching
	db *gorm.DB
	// apiKey to access quandl
	apiKey string
	// Map from the ticker symbol to the last sucessful refresh.
	ref map[string]time.Time
	// protects the access to the map.
	refGuard sync.RWMutex
}

// NewCache creates a new file-based cache, locally.
func NewCache(apiKey string) *Cache {
	return newCache(apiKey, "mystock.db")
}

// NewMemoryCache creates a Cache in memory.
// Mainly used for testing.
func NewMemoryCache(apiKey string) *Cache {
	return newCache(apiKey, ":memory:")
}

// newCache actually creates and initialize the cache.
func newCache(apiKey, fname string) *Cache {
	var err error

	c := new(Cache)
	c.apiKey = apiKey
	c.db, err = gorm.Open("sqlite3", fname)
	if err != nil {
		panic(err)
	}
	err = c.db.AutoMigrate(&quandl.Record{}, &ref{}).Error
	if err != nil {
		panic(err)
	}
	c.ref = make(map[string]time.Time)

	// Restore saved refs if any from the db.
	// That will avoid unnecessary refresh, and identify all Tickers.
	c.restoreRefs()

	return c
}

// Close the underlying database.
// Required to flush the cache when file.
func (c *Cache) Close() {
	// save refs for the next time.
	c.saveRefs()
	c.db.Close()
}

// Size provides count of total records in cache.
func (c *Cache) Size() (n int) {
	err := c.db.Model(&quandl.Record{}).Count(&n).Error
	if err != nil {
		panic(err)
	}
	return n
}

// Dump dumps the data base.
// Used for testing/debugging.
func (c *Cache) Dump() {
	var rr []quandl.Record
	err := c.db.Find(&rr).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are  %d records in dbase\n", len(rr))
	for i, r := range rr {
		fmt.Println(r)
		if i >= 10 {
			fmt.Println("... truncated ...")
			break
		}
	}
	fmt.Println()
}

// refresh refresh the cache for the provided record.
// If cache was still fresh, do nothing, return false.
// Retun true if actual refresh happened.
// Thead safe.
func (c *Cache) refresh(ticker string) bool {

	if len(ticker) == 0 {
		panic("cannot refresh an empty string ticker")
	}
	// check if refresh is needed ...
	// at least 6 hour between refreshes.
	c.refGuard.RLock()
	last, ok := c.ref[ticker]
	if ok && last.Add(6*time.Hour).After(time.Now()) {
		// no refresh needed
		c.refGuard.RUnlock()
		return false
	}

	// change for rw lock
	c.refGuard.RUnlock()
	c.refGuard.Lock()

	// debug
	fmt.Println("Refreshing ", ticker)

	// do the actual refresh
	// refresh locked during this refresh.
	quandl.New(c.apiKey, "EURONEXT").WalkDataset(ticker, c.saveRecord)
	c.ref[ticker] = time.Now()

	// release lock
	c.refGuard.Unlock()
	return true
}

// save record in database, updating if needed.
func (c *Cache) saveRecord(r quandl.Record) {
	err := c.db.Save(&r).Error
	if err != nil {
		panic(err)
	}
}

// ListTickers will list the tickers managed in the cache in a thread safe way.
func (c *Cache) ListTickers() []string {

	c.refGuard.RLock()

	res := make([]string, 0, len(c.ref))
	for s := range c.ref {
		res = append(res, s)
	}

	c.refGuard.RUnlock()

	return res
}
