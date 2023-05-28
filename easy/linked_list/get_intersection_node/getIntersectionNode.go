// https://leetcode.com/problems/intersection-of-two-linked-lists/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vals := make(map[*ListNode]struct{})

	for headA != nil {
		vals[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := vals[headB]; ok {
			return headB
		} else {
			vals[headB] = struct{}{}
		}
		headB = headB.Next
	}

	return nil

}
