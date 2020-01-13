package qdl

import "testing"

import "os"

import "time"

func TestMemoryDBCreation(t *testing.T) {
	c := NewMemoryDB()
	defer c.Close()
}

func TestFileCacheDBCreation(t *testing.T) {
	defer os.Remove("test.db")
	c := newCacheDB("test.db")
	defer c.Close()
}

func TestUpdateRecord(t *testing.T) {
	c := NewMemoryDB()
	defer c.Close()
	q := NewQDL().SetStartDate(time.Now().Add(-time.Hour * 24 * 5))
	c.Refresh(q, "ML")
	c.Dump()
	c1, _ := c.Count()
	c.Refresh(q, "ML")
	c2, _ := c.Count()
	if c1 != c2 {
		t.Fatal("Count should not have changed - updates expected to be idempotent")
	}
}
