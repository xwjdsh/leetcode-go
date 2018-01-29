package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestThreeSum(t *testing.T) {
	result1Map := map[string]bool{
		"0,-1,1": true,
		"0,1,-1": true,
		"-1,0,1": true,
		"-1,1,0": true,
		"1,0,-1": true,
		"1,-1,0": true,
	}
	result2Map := map[string]bool{
		"-1,-1,2": true,
		"-1,2,-1": true,
		"2,-1,-1": true,
	}

	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	if len(result) != 2 {
		t.Error("test threeSum error")
	}
	rs := result[0]
	strs := []string{}
	for _, r := range rs {
		strs = append(strs, strconv.Itoa(r))
	}
	result1Str := strings.Join(strs, ",")

	rs = result[1]
	strs = []string{}
	for _, r := range rs {
		strs = append(strs, strconv.Itoa(r))
	}
	result2Str := strings.Join(strs, ",")

	if result1Map[result1Str] {
		result1Map = map[string]bool{}
	} else if result2Map[result1Str] {
		result2Map = map[string]bool{}
	} else {
		t.Error("test threeSum error")
	}

	if !result1Map[result2Str] && !result2Map[result2Str] {
		t.Error("test threeSum error")
	}
}
