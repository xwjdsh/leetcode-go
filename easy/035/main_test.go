package main

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	data := []int{1, 3, 5, 6}
	if searchInsert(data, 5) != 2 || searchInsert(data, 2) != 1 || searchInsert(data, 7) != 4 || searchInsert(data, 0) != 0 {
		t.Error("searchInsert test failed")
	}
}
