package cache

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/xavier268/mystock/configuration"
	"github.com/xavier268/mystock/quandl"

	// blank import used for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Cache is the DB access object used to cach or retrieve data.
type Cache struct {
	// database for caching
	*gorm.DB
	// apiKey to access quandl
	apiKey string
	// Map from the ticker symbol to the last sucessful refresh.
	ref map[string]time.Time
	// protects the access to the map.
	refGuard sync.RWMutex
}

// NewCache creates a new file-based cache, locally.
func NewCache(conf configuration.Conf) *Cache {
	return newCache(conf, "mystock.db")
}

// NewMemoryCache creates a Cache in memory.
// Mainly used for testing.
func NewMemoryCache(conf configuration.Conf) *Cache {
	return newCache(conf, ":memory:")
}

// newCache actually creates and initialize the cache.
func newCache(conf configuration.Conf, fname string) *Cache {
	var err error

	c := new(Cache)
	c.apiKey = conf.APISecretKey
	c.DB, err = gorm.Open("sqlite3", fname)
	if err != nil {
		panic(err)
	}
	err = c.DB.AutoMigrate(&quandl.Record{}, &ref{}).Error
	if err != nil {
		panic(err)
	}
	c.ref = make(map[string]time.Time)

	// Restore saved refs if any from the db.
	// That will avoid unnecessary refresh, and identify all Tickers.
	// Extract tickers from config file as default tickers.
	c.restoreRefs(conf.Tickers()...)

	return c
}

// Close the underlying database.
// Required to flush the cache when file.
func (c *Cache) Close() {
	// save refs for the next time.
	c.saveRefs()
	c.DB.Close()
}

// Size provides count of total records in cache.
func (c *Cache) Size() (n int) {
	err := c.DB.Model(&quandl.Record{}).Count(&n).Error
	if err != nil {
		panic(err)
	}
	return n
}

// Dump dumps the data base.
// Used for testing/debugging.
func (c *Cache) Dump() {
	var rr []quandl.Record
	var rfs []ref

	// dump records
	err := c.DB.Find(&rr).Error
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

	// dump refs
	err = c.DB.Find(&rfs).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are  %d refs in dbase\n", len(rfs))
	for i, r := range rfs {
		fmt.Println(r)
		if i >= 10 {
			fmt.Println("... truncated ...")
			break
		}
	}
	fmt.Println()
}

// Refresh the cache, if needed, for all tickers.
func (c *Cache) Refresh() {

	lt := c.ListTickers()
	for _, t := range lt {
		c.refresh(t)
	}
}

// refresh refresh the cache for the provided record.
// If cache was still fresh, do nothing, return false.
// Retun true if actual refresh happened.
// Thead safe.
func (c *Cache) refresh(ticker string) bool {

	if len(ticker) == 0 {
		panic("cannot refresh an empty string ticker")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Could not refresh cache : ", r)
		}
	}()

	// check if refresh is needed ...
	// at least 6 hour between refreshes.
	c.refGuard.RLock()
	last, ok := c.ref[ticker]
	if ok && (last != time.Time{}) && last.Add(6*time.Hour).After(time.Now()) {
		// no refresh needed
		fmt.Println("No refresh needed for " + ticker)
		c.refGuard.RUnlock()
		return false
	}
	c.refGuard.RUnlock()
	fmt.Println("Preparing to refresh " + ticker)

	// Query most recent ticker date,
	// with a few days to be sure ...
	since := c.MostRecent(ticker)
	if (since != time.Time{}) {
		since = since.Add(-3 * 24 * time.Hour)
	}
	fmt.Println("Refresh needed since " + since.String())

	// Set rw lock
	c.refGuard.Lock()
	defer c.refGuard.Unlock()
	// debug
	fmt.Println("Refreshing ", ticker)

	// do the actual refresh
	// refresh locked during this refresh.
	quandl.New(c.apiKey, "EURONEXT", quandl.OptionStartDate(since)).
		WalkDataset(ticker, c.saveRecord)
	c.ref[ticker] = time.Now()
	return true
}

// save record in database, updating if needed.
func (c *Cache) saveRecord(r quandl.Record) {
	err := c.DB.Save(&r).Error
	if err != nil {
		panic(err)
	}
}

// ListTickers will list the tickers managed
// in the cache in a thread safe way.
func (c *Cache) ListTickers() []string {

	c.refGuard.RLock()

	res := make([]string, 0, len(c.ref))
	for s := range c.ref {
		res = append(res, s)
	}

	c.refGuard.RUnlock()

	return res
}

// MostRecent is the most recent date for that ticker in the database.
func (c *Cache) MostRecent(ticker string) time.Time {

	// If ticker is not known, return zero value time.
	if _, ok := c.ref[ticker]; !ok {
		return time.Time{}
	}

	var rr []quandl.Record
	err := c.DB.Model(&quandl.Record{}).
		Where("serie = ?", strings.ToUpper(ticker)).
		Order("date DESC").
		Limit(1).
		Find(&rr).Error
	if err != nil {
		panic(err)
	}
	if len(rr) == 0 {
		return time.Time{}
	}
	return rr[0].Date
}
