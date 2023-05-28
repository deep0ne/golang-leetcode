// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	head = head.Next
	ans := prev

	for head != nil {
		if head.Val == prev.Val {
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}

	return ans
}
