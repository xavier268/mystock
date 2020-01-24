package monitor

import (
	"fmt"
	"testing"
)

func TestCheckPrice(t *testing.T) {
	for tt := range m.LoadPortfolio() {
		fmt.Println("Testing price change for ", tt)
		mess, e := CheckPriceChange("MC", 0.)(m)
		if e != nil {
			t.Fatal(e)
		}
		if len(mess) == 0 {
			t.Fatal("Unexpected empty message ?!")
		}
	}
}

func TestCheckPriceAll(t *testing.T) {

	fmt.Println("Testing price change ALL ")
	mess, e := CheckPriceChangeAll(0.)(m)
	if e != nil {
		t.Fatal(e)
	}
	if len(mess) == 0 {
		t.Fatal("Unexpected empty message ?!")
	}

}

func TestCheckGainLoss(t *testing.T) {
	mess, e := CheckGainLoss(0.)(m)
	if e != nil {
		t.Fatal(e)
	}
	if len(mess) == 0 {
		t.Fatal("Unexpected empty message")
	}
}
