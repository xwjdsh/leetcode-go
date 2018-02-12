package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for pre, next := head, head.Next; pre != nil && next != nil; {
		if pre.Val == next.Val {
			pre.Next = nil
			next = next.Next
		} else {
			pre.Next = next
			pre, next = next, next.Next
		}
	}
	return head
}
