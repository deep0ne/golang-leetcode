// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/?envType=study-plan-v2&envId=top-interview-150

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	prev := dummy
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}
		head = head.Next
	}
	return dummy.Next
}
