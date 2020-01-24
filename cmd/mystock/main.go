package main

import (
	"fmt"

	"github.com/xavier268/mystock/monitor"
)

func main() {
	fmt.Println("Running my stock monitoring utility")
	percent := 1 / 100.
	fmt.Printf("Threshold is set at : %.1f%%\n", 100*percent)
	m := monitor.NewMonitor(monitor.AlertLog(),
		monitor.CheckGainLoss(percent),
		monitor.CheckPriceChangeAll(percent),
	)
	defer m.Close()

	// Run all checks once
	m.CheckAll()
}
