package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	if !isValid("()") || !isValid("()[]{}") || isValid("(]") || isValid("([)]") {
		t.Error("isValid test failed")
	}
}
