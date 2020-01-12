package qdl

import "testing"

import "os"

func TestMemoryDBCreation(t *testing.T) {
	c := NewMemoryDB()
	defer c.Close()
}

func TestFileCacheDBCreation(t *testing.T) {
	defer os.Remove("test.db")
	c := newCacheDB("test.db")
	defer c.Close()
}
