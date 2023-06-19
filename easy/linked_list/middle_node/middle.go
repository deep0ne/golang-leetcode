// https://leetcode.com/problems/middle-of-the-linked-list/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	curr, index := head, 0
	for curr != nil {
		curr = curr.Next
		index++
	}
	mid := 0
	for head != nil && mid != index/2 {
		head = head.Next
		mid++
	}
	return head
}
