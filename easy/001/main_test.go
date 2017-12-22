package main

import (
	"reflect"
	"testing"
)

var (
	nums   = []int{2, 7, 11, 15}
	target = 9

	expect_1 = []int{0, 1}
	expect_2 = []int{1, 0}
)

func TestTwoSum_1(t *testing.T) {
	resul := twoSum_1(nums, target)
	if !(reflect.DeepEqual(resul, expect_1) || reflect.DeepEqual(resul, expect_2)) {
		t.Error("twoSum_1 test failed")
	}
}

func TestTwoSum_2(t *testing.T) {
	resul := twoSum_2(nums, target)
	if !(reflect.DeepEqual(resul, expect_1) || reflect.DeepEqual(resul, expect_2)) {
		t.Error("twoSum_2 test failed")
	}
}

func TestTwoSum(t *testing.T) {
	resul := twoSum(nums, target)
	if !(reflect.DeepEqual(resul, expect_1) || reflect.DeepEqual(resul, expect_2)) {
		t.Error("twoSum test failed")
	}
}
