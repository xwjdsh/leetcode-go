package main

import (
	"testing"
)

func TestHammingDistance_1(t *testing.T) {
	if hammingDistance_1(1, 4) != 2 {
		t.Error("hammingDistance_1 test failed")
	}
}

func TestHammingDistance(t *testing.T) {
	if hammingDistance(1, 4) != 2 {
		t.Error("hammingDistance test failed")
	}
}
