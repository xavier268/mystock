package qdl

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// blank import used for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// CacheDB is the DB access object used to cach or retrieve data.
type CacheDB struct {
	db *gorm.DB
}

// NewFileDB creates a new db file locally.
func NewFileDB() *CacheDB {
	return newCacheDB("mystock.db")
}

// NewMemoryDB creates a cacheDB in memory.
// Mainly used for testing.
func NewMemoryDB() *CacheDB {
	return newCacheDB(":memory:")
}

func newCacheDB(fname string) *CacheDB {
	var err error

	c := new(CacheDB)
	c.db, err = gorm.Open("sqlite3", fname)
	if err != nil {
		panic(err)
	}
	err = c.db.AutoMigrate(&Record{}).Error
	if err != nil {
		panic(err)
	}
	return c
}

// Close the database.
// Required after creating a new one.
func (c *CacheDB) Close() {
	c.db.Close()
}

// Update record.
func (c *CacheDB) Update(r *Record) {
	// fmt.Println("Updating : ", r)
	err := c.db.Save(r).Error
	if err != nil {
		panic(err)
	}
}

// Refresh the history for the provided tickers.
func (c *CacheDB) Refresh(q *QDL, tickers ...string) {
	if len(tickers) == 0 {
		return
	}
	for _, t := range tickers {
		q.Refresh(Code{"EURONEXT", t}, c.Update)
	}
}

// Dump dumps the data base. Used for testing/debugging.
func (c *CacheDB) Dump() {
	var rr []Record
	err := c.db.Find(&rr).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(len(rr), " records in dbase\n", rr)
}
