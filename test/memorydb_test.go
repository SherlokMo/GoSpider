package test

import (
	"goSpider/memorydb"
	"testing"
)

func TestInsertion(t *testing.T) {
	memdb := memorydb.NewMemorydb[string, int]()

	memdb.Store("A", 1)
	memdb.Store("B", 2)

	if OK := memdb.Exist("A"); !OK {
		t.Error("Expected True recived false")
	}

	if OK := memdb.Exist("C"); OK {
		t.Error("Expected True recived false")
	}
}
