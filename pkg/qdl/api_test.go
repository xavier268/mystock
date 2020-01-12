package qdl

import (
	"fmt"
	"testing"
)

func TestMakeURL(t *testing.T) {

	q := New()
	u := q.makeURL("euronext", "hello")
	// fmt.Println(u.String())
	if u.String() != "https://www.quandl.com/api/v3/datasets/EURONEXT/HELLO/data.json?api_key="+APISecretKey {
		t.Fatal(u.String())
	}

}

func TestGetCode1(t *testing.T) {
	q := New()

	// Valid - should not panic
	h := q.GetHistory("euronext", "adyen")

	fmt.Println("\n========= HISTORY ==============\n", h)

}

func TestGetCode2(t *testing.T) {
	// Invalid - should panic
	defer expectPanic(t)

	q := New()

	q.GetHistory("euronext", "testtsttest")

}

// ***************** utilities *****************
func expectPanic(t *testing.T) {
	fmt.Println("Panic recovery ...")
	if r := recover(); r != nil {
		fmt.Println("Panic was expected, ok !")
	} else {
		t.Fatal("Code should have panicked !?")
	}
}
