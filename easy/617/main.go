package main

import (
	"strconv"
	"strings"
)

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	result := []string{}
	if tn == nil {
		result = append(result, "nil")
	} else {
		result = append(result, strconv.Itoa(tn.Val))
		if tn.Left != nil {
			result = append(result, tn.Left.String())
		}
		if tn.Right != nil {
			result = append(result, tn.Right.String())
		}
	}
	return strings.Join(result, ",")
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	result := &TreeNode{}
	var t1Left, t2Left, t1Right, t2Right *TreeNode
	if t1 != nil {
		result.Val += t1.Val
		t1Left, t1Right = t1.Left, t1.Right
	}
	if t2 != nil {
		result.Val += t2.Val
		t2Left, t2Right = t2.Left, t2.Right
	}
	if t1Left != nil || t2Left != nil {
		result.Left = mergeTrees(t1Left, t2Left)
	}
	if t1Right != nil || t2Right != nil {
		result.Right = mergeTrees(t1Right, t2Right)
	}
	return result
}
