// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums) / 2
	root := &TreeNode{Val: nums[n]}
	root.Left = helper(nums[:n])
	root.Right = helper(nums[n+1:])
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return helper(nums)
}
