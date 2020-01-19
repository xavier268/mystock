package configuration

import (
	"fmt"
	"testing"
)

func TestConfiguration(t *testing.T) {
	c := Load()
	fmt.Println("Configuration read :", c)

	tl := c.Tickers()
	fmt.Printf("Tickers in configuration file : %+v\n", tl)
}
