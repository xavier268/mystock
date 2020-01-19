package cache

import (
	"fmt"
	"time"
)

// ref saves/restore the refresh map from the db.
type ref struct {
	Ticker string `gorm:"primary_key"`
	When   time.Time
}

// Save ref to database, before closing.
func (c *Cache) saveRefs() {

	c.refGuard.RLock()
	defer c.refGuard.RUnlock()

	for r := range c.ref {
		ref := ref{r, c.ref[r]}
		err := c.DB.Save(ref).Error
		if err != nil {
			panic(err)
		}
	}

}

// Restore refs from database, during init.
// A list fo initial tickers can be provided, that will be set
// as not fresh.
func (c *Cache) restoreRefs(confTickers ...string) {
	var refs []ref
	err := c.DB.Model(&ref{}).Find(&refs).Error
	if err != nil {
		fmt.Println("could not restore ref map ?!")
		panic(err)
	}
	c.refGuard.Lock()
	// Load predefined default tickers
	for _, t := range confTickers {
		c.ref[t] = time.Time{}
	}
	// Restore Tickers from database
	for _, r := range refs {
		c.ref[r.Ticker] = r.When
	}
	c.refGuard.Unlock()
}
