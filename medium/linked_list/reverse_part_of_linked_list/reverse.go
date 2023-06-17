// https://leetcode.com/problems/reverse-linked-list-ii/description/?envType=study-plan-v2&envId=top-interview-150
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := new(ListNode)
	head, dummy.Next = dummy, head
	for i := 0; i < left-1; i++ {
		head = head.Next
	}

	var curr, prev *ListNode = head.Next, nil
	for i := 0; i < right-left+1; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head.Next.Next = curr
	head.Next = prev
	return dummy.Next
}
