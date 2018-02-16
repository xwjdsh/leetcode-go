package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	result := &ListNode{
		Val: head.Val,
	}

	for head.Next != nil {
		tmp := &ListNode{
			Val: head.Next.Val,
		}
		tmp.Next = result
		result = tmp
		head = head.Next
	}
	return result
}
