// https://leetcode.com/problems/remove-linked-list-elements/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	prev := head
	ans := prev
	head = head.Next
	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}
	if ans != nil && ans.Val == val {
		return ans.Next
	}
	return ans
}
