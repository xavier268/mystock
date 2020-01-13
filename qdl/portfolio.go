package qdl

// CreateOrder save the Order in the Portfolio.
func (c *CacheDB) CreateOrder(o *Order) {
	err := c.db.Create(o).Error
	if err != nil {
		panic(err)
	}
}
