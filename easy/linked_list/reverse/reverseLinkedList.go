// https://leetcode.com/problems/reverse-linked-list/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		curr := head.Next
		head.Next = prev
		prev = head
		head = curr
	}
	return prev
}
