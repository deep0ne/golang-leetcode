// https://leetcode.com/problems/palindrome-linked-list/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	memory := make([]int, 0)
	for head != nil {
		memory = append(memory, head.Val)
		head = head.Next
	}
	for i, j := 0, len(memory)-1; i < len(memory)/2; i, j = i+1, j-1 {
		if memory[i] != memory[j] {
			return false
		}
	}
	return true
}
