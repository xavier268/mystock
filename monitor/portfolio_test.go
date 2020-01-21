package monitor

import (
	"fmt"
	"testing"
)

func TestLastClosingPrice(t *testing.T) {
	p := m.LastClosingPrice("ML")
	fmt.Println("Last closing ML price : ", p)
}
