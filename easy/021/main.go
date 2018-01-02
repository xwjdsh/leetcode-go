package main

import "strconv"

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	var result string
	for ; ln != nil; ln = ln.Next {
		result += strconv.Itoa(ln.Val)
	}
	return result
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	result := &ListNode{}
	if l1 != nil && l2 != nil {
		result.Val = l1.Val
		if l1.Val < l2.Val {
			result.Val = l1.Val
			l1 = l1.Next
		} else {
			result.Val = l2.Val
			l2 = l2.Next
		}
	} else if l1 != nil {
		result.Val = l1.Val
		l1 = l1.Next
	} else {
		result.Val = l2.Val
		l2 = l2.Next
	}

	if l1 != nil || l2 != nil {
		result.Next = mergeTwoLists(l1, l2)
	}
	return result
}
