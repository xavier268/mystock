package monitor

import "testing"

func TestAlertLog(t *testing.T) {

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
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlertSNS(t *testing.T) {

	c := loadConfiguration()
	err := AlertSNS(c.SNSTopic)("Message SMS de test")
	if err != nil {
		t.Fatal(err)
	}
}
