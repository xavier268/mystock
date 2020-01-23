package monitor

import (
	"fmt"
	"os"
	"testing"
)

var m *Monitor

func TestMain(t *testing.M) {
	m = NewMonitor(nil)
	// DEBUG : Setting logmode will display actual sql
	// m.cache.DB.LogMode(true)
	fmt.Println("Test Monitor constructed : ")
	m.cache.Dump()
	e := t.Run()
	m.Close() // Don't defer, exit might not call it !
	os.Exit(e)
}
