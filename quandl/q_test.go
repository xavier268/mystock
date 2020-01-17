package quandl

import (
	"fmt"
	"testing"
)

func TestAPISecretKeyConfigured(t *testing.T) {
	// This will fail if secret key is not configured.
	// Make it available in secret_test.go
	fmt.Println("Test secret key used : ", mysecret)
}

func TestConstructQ(t *testing.T) {
	New(mysecret, "euroNExt").WalkDataset("ML", doNothing)
}

func TestConstructBadSerie(t *testing.T) {
	defer expectPanic(t)
	New(mysecret, "euroNExt").WalkDataset("wrongDataSerie", doNothing)
}

func TestConstructBadSource(t *testing.T) {
	defer expectPanic(t)
	New(mysecret, "badSource").WalkDataset("ML", doNothing)
}

// This test is skipped, because, surpirsingly,
// the web service does not seem the check the API KEY,
// and appears to accepts any string ?!
func TestConstructBadAPIKey(t *testing.T) {

	t.Skip()

	defer expectPanic(t)
	New("testKeyIsNotValid", "EURONEXT").WalkDataset("ML", doNothing)
}
