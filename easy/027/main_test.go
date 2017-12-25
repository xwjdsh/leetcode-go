package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement_1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	resul := 2
	expect := []int{2, 2}

	if removeElement_1(nums, val) != resul || !reflect.DeepEqual(nums[:resul], expect) {
		t.Error("removeElement_1 test failed")
	}
}

func TestRemoveElement_2(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	resul := 2
	expect := []int{2, 2}

	if removeElement_2(nums, val) != resul || !reflect.DeepEqual(nums[:resul], expect) {
		t.Error("removeElement_2 test failed")
	}
}

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	resul := 2
	expect := []int{2, 2}

	if removeElement(nums, val) != resul || !reflect.DeepEqual(nums[:resul], expect) {
		t.Error("removeElement test failed")
	}
}
