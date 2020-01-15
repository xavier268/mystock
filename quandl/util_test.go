package quandl

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer expectPanic(t)
	panic("test")
}

func expectPanic(t *testing.T) {
	if r := recover(); r != nil {
		fmt.Println("Panic was expected, all is good !")
	} else {
		t.Fatal("Code should have panicked !?")
	}
}

var doNothing RecordProcessor = func(r Record) {}
