package main

import (
	"testing"
)

func TestJudgeCircle(t *testing.T) {
	if !judgeCircle("UD") || judgeCircle("LL") {
		t.Error("judgeCircle test failed")
	}
}
