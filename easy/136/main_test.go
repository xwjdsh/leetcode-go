package main

import (
	"testing"
)

func TestSingleNumber_1(t *testing.T) {
	if singleNumber_1([]int{1, 2, 3, 3, 2}) != 1 {
		t.Error("singleNumber_1 test failed")
	}
}

func TestSingleNumber(t *testing.T) {
	if singleNumber([]int{1, 2, 3, 3, 2}) != 1 {
		t.Error("singleNumber test failed")
	}
}
