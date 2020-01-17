package monitor

import (
	"fmt"
	"testing"
)

func TestConfiguration(t *testing.T) {
	c := loadConfiguration()
	fmt.Println(c)
}
