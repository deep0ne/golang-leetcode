// https://leetcode.com/problems/add-two-numbers/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	memory, dummy := 0, new(ListNode)
	for node := dummy; l1 != nil || l2 != nil || memory > 0; node = node.Next {
		if l1 != nil {
			memory += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			memory += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{memory % 10, nil}
		memory /= 10
	}
	return dummy.Next
}
