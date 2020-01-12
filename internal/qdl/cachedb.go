package qdl

import (
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
	c.db.AutoMigrate(&Record{})
	return c
}

// Close the database.
// Required after creating a new one.
func (c *CacheDB) Close() {
	c.db.Close()
}
