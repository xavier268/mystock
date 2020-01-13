package qdl

// CreateOrder save the Order in the Portfolio.
func (c *CacheDB) CreateOrder(o *Order) {
	err := c.db.Create(o).Error
	if err != nil {
		panic(err)
	}
}

// PortfolioHistValue computes the historiacl value of the portfolio.
func (c *CacheDB) PortfolioHistValue() float64 {
	var res struct{ Total float64 }
	c.db.Model(&Order{}).Select("SUM( price * volume ) as total ").Scan(&res)
	//fmt.Println(res)
	return res.Total
}
