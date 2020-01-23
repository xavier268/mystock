package main

import (
	"fmt"

	"github.com/xavier268/mystock/monitor"
)

func main() {
	fmt.Println("Running my stock monitoring utility")
	percent := 0.1 / 100.
	fmt.Printf("Theshold is set at : %.1f%%\n", 100*percent)
	m := monitor.NewMonitor(monitor.AlertLog(),
		monitor.CheckGainLoss(percent),
		monitor.CheckPriceChange("AIR", percent),
		monitor.CheckPriceChange("SAN", percent),
		monitor.CheckPriceChange("MC", percent),
		monitor.CheckPriceChange("ML", percent),
	)
	defer m.Close()

	// Run all checks once
	m.CheckAll()
}
