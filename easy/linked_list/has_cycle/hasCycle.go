package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	m := make(map[*ListNode]struct{})
	m[head] = struct{}{}
	for head.Next != nil {
		next := head.Next
		if _, ok := m[next]; ok {
			return true
		} else {
			m[next] = struct{}{}
		}
		head = head.Next
	}
	return false
}
