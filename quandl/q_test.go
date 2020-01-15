package quandl

import "testing"

func TestAPISecretKeyConfigured(t *testing.T) {
	// This will fail is secret key is not configured.
	APISecretKey()
}

func TestConstructQ(t *testing.T) {
	New("euroNExt").WalkDataset("ML", doNothing)
}

func TestPanic(t *testing.T) {
	defer expectPanic(t)
	panic("test")
}

func TestConstructBadTicker(t *testing.T) {
	defer expectPanic(t)
	New("euroNExt").WalkDataset("wrongTicker", doNothing)
}

func TestConstructBadSource(t *testing.T) {
	defer expectPanic(t)
	New("badSource").WalkDataset("ML", doNothing)
}
