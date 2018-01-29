package main

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {
	p1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	q1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	p2 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	q2 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}

	p3 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}
	q3 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}

	if !isSameTree(p1, q1) || isSameTree(p2, q2) || isSameTree(p3, q3) {
		t.Error("isSameTree test failed")
	}
}
