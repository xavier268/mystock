package monitor

import (
	"os"
	"testing"
)

var m *Monitor

func TestMain(t *testing.M) {
	m = NewMonitor(nil)
	// DEBUG : Setting logmode will display actual sql
	// m.cache.DB.LogMode(true)
	e := t.Run()
	m.Close()
	os.Exit(e)
}
