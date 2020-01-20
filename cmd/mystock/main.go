package main

import (
	"fmt"

	"github.com/xavier268/mystock/monitor"
)

func main() {
	fmt.Println("Running my stock monitoring utility")
	// conf := configuration.Load()
	m := monitor.NewMonitor(monitor.AlertLog())
	defer m.Close()

}
