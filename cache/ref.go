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
	for r := range c.ref {
		ref := ref{r, c.ref[r]}
		err := c.db.Save(ref).Error
		if err != nil {
			panic(err)
		}
	}
	c.refGuard.RUnlock()
}

// Restore refs from database, during init.
func (c *Cache) restoreRefs() {
	var refs []ref
	err := c.db.Model(&ref{}).Find(&refs).Error
	if err != nil {
		fmt.Println("could not restore ref map ?!")
		panic(err)
	}
	c.refGuard.Lock()
	for _, r := range refs {
		c.ref[r.Ticker] = r.When
	}
	c.refGuard.Unlock()
}
