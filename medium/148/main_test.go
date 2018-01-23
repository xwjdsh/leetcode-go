package main

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	l := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	result := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	if !reflect.DeepEqual(*sortList(l), *result) {
		t.Error("sortList test failed")
	}
}
