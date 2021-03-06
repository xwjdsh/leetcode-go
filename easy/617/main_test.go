package main

import (
	"testing"
)

func TestMergeTrees(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	t2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	if mergeTrees(t1, t2).String() != "3,4,5,4,5,7" {
		t.Error("mergeTrees test failed")
	}
}
