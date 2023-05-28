// https://leetcode.com/problems/merge-two-sorted-lists/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newListHead := &ListNode{}

	ans := newListHead

	for list1 != nil || list2 != nil {
		if list1 == nil {
			newListHead.Next = list2
			newListHead = newListHead.Next
			list2 = list2.Next
			continue
		}

		if list2 == nil {
			newListHead.Next = list1
			newListHead = newListHead.Next
			list1 = list1.Next
			continue
		}

		if list1.Val < list2.Val {
			newListHead.Next = list1
			list1 = list1.Next
		} else {
			newListHead.Next = list2
			list2 = list2.Next
		}

		newListHead = newListHead.Next
	}

	return ans.Next
}
