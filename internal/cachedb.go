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

// New creates a new db file locally.
func (c *CacheDB) New() *CacheDB {
	var err error

	c = new(CacheDB)
	c.db, err = gorm.Open("sqlite3", "mystock.db")
	if err != nil {
		panic(err)
	}
	return c

}

// Close the database.
// Required after creating a new one.
func (c *CacheDB) Close() {
	c.Close()
}
