// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums) / 2
	root := &TreeNode{Val: nums[n]}
	root.Left = sortedArrayToBST(nums[:n])
	root.Right = sortedArrayToBST(nums[n+1:])
	return root
}
