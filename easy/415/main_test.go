package main

import (
	"testing"
)

func TestAddStrings(t *testing.T) {
	if addStrings("111", "222") != "333" {
		t.Error("addStrings test failed")
	}
}
