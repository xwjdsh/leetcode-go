package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if reverse(123) != 321 || reverse(-123) != -321 || reverse(120) != 21 {
		t.Error("reverse test failed")
	}
}
