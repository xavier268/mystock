package quandl

import "testing"

func TestAPISecretKeyConfigured(t *testing.T) {
	// This will fail is secret key is not configured.
	APISecretKey()
}

func TestConstructQ(t *testing.T) {
	New("euroNExt").WalkDataset("ML", doNothing)
}

func TestConstructBadSerie(t *testing.T) {
	defer expectPanic(t)
	New("euroNExt").WalkDataset("wrongDataSerie", doNothing)
}

func TestConstructBadSource(t *testing.T) {
	defer expectPanic(t)
	New("badSource").WalkDataset("ML", doNothing)
}
