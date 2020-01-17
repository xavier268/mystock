package monitor

import "testing"

func TestAlerts(t *testing.T) {
	var err error
	err = AlertLog()("")
	if err != nil {
		t.Fatal(err)
	}
	err = AlertLog()("testing AlertLog")
	if err != nil {
		t.Fatal(err)
	}
	err = AlertLog()(nil)

}
