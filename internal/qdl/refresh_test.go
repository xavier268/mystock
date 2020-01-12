package qdl

import (
	"fmt"
	"testing"
)

func TestMakeURL(t *testing.T) {

	q := NewQDL()
	u := q.makeURLString(Code{"euronext", "hello"})
	// fmt.Println(u.String())
	if u != "https://www.quandl.com/api/v3/datasets/EURONEXT/HELLO/data.json?api_key="+APISecretKey {
		t.Fatal(u)
	}

}

func TestGetCode1(t *testing.T) {
	q := NewQDL()

	// Valid - should not panic, but display records
	q.Refresh(Code{"euronext", "adyen"}, doNothing)

}

func TestGetCode2(t *testing.T) {
	// Invalid - should panic
	defer expectPanic(t)

	q := NewQDL()

	q.Refresh(Code{"euronext", "testetst"}, printRecord)

}

// =============================================
//               test utilities
// =============================================

func expectPanic(t *testing.T) {
	fmt.Println("Panic recovery ...")
	if r := recover(); r != nil {
		fmt.Println("Panic was expected, ok !")
	} else {
		t.Fatal("Code should have panicked !?")
	}
}

func printRecord(r Record) {
	fmt.Println(r)
}

func doNothing(r Record) {
}
