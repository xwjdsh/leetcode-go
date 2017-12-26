package main

import (
	"reflect"
	"testing"
)

var (
	nums = []int{4, 3, 2, 7, 8, 2, 3, 1}

	expect_1 = []int{2, 3}
	expect_2 = []int{3, 2}
)

func TestFindDuplicates_1(t *testing.T) {
	result := findDuplicates_1(nums)
	if !(reflect.DeepEqual(result, expect_1) || reflect.DeepEqual(result, expect_2)) {
		t.Error("findDuplicates_1 test failed")
	}
}

func TestFindDuplicates(t *testing.T) {
	result := findDuplicates(nums)
	if !(reflect.DeepEqual(result, expect_1) || reflect.DeepEqual(result, expect_2)) {
		t.Error("findDuplicates test failed")
	}
}
