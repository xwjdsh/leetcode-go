package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for a := head; a.Next != nil; a = a.Next {
		for b := a.Next; b != nil; b = b.Next {
			if a.Val > b.Val {
				a.Val, b.Val = b.Val, a.Val
			}
		}
	}
	return head
}
