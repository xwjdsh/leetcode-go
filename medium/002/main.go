package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	for extra, pre := 0, result; l1 != nil || l2 != nil || extra > 0; {
		total := extra
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}
		node := &ListNode{Val: total % 10}
		extra = total / 10

		if pre == nil {
			result = node
		} else {
			pre.Next = node
		}
		pre = node
	}
	return result
}
