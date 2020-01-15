package quandl

import (
	"fmt"
	"testing"
)

func expectPanic(t *testing.T) {
	fmt.Println("Panic recovery ...")
	if r := recover(); r != nil {
		fmt.Println("Panic was expected, ok !")
	} else {
		t.Fatal("Code should have panicked !?")
	}
}

var doNothing RecordProcessor = func(r Record) {}
