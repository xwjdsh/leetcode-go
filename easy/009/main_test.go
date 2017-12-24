package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(1221) || isPalindrome(1000021) {
		t.Error("isPalindrome test failed")
	}
}
