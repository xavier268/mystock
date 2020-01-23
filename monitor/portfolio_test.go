package monitor

import (
	"fmt"
	"testing"

	"github.com/xavier268/mystock/configuration"
)

func TestLastClosingPrice(t *testing.T) {
	p := m.LastClosingPrice("ML")
	fmt.Println("Last closing ML price : ", p)
}

func TestPortfolio(t *testing.T) {
	p := m.LoadPortfolio()
	p.Dump()
	h, l := p.Values()
	fmt.Println("Portfolio values : ", h, l)
	if h == 0 || l == 0 {
		fmt.Println("Configuration : ", configuration.Load())
		p.Dump()
		t.Fatal("Unexpected porfolio values : ", h, l)
	}

}
