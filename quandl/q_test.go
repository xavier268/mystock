package quandl

import (
	"fmt"
	"testing"
	"time"

	"github.com/xavier268/mystock/configuration"
)

func TestAPISecretKeyConfigured(t *testing.T) {
	// This will fail if secret key is not configured.
	// Make it available in secret_test.go
	fmt.Println("Test secret key used : ", configuration.Load().APISecretKey)
}

func TestConstructQ(t *testing.T) {
	New(configuration.Load().APISecretKey, "euroNExt").WalkDataset("ML", doNothing)
}
func TestConstructQWithSince(t *testing.T) {
	New(configuration.Load().APISecretKey, "euroNExt", OptionStartDate(time.Now())).WalkDataset("ML", doNothing)
}

func TestConstructBadSerie(t *testing.T) {
	defer expectPanic(t)
	New(configuration.Load().APISecretKey, "euroNExt").WalkDataset("wrongDataSerie", doNothing)
}

func TestConstructBadSource(t *testing.T) {
	defer expectPanic(t)
	New(configuration.Load().APISecretKey, "badSource").WalkDataset("ML", doNothing)
}

// This test is skipped, because, surpirsingly,
// the web service does not seem the check the API KEY,
// and appears to accepts any string ?!
func TestConstructBadAPIKey(t *testing.T) {

	t.Skip()

	defer expectPanic(t)
	New("testKeyIsNotValid", "EURONEXT").WalkDataset("ML", doNothing)
}
