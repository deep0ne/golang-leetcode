// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, nums []int) []int {

	if root.Left != nil {
		nums = helper(root.Left, nums)
	}
	nums = append(nums, root.Val)
	if root.Right != nil {
		nums = helper(root.Right, nums)
	}
	return nums
}

func kthSmallest(root *TreeNode, k int) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	nums := helper(root, []int{})
	fmt.Println(nums)
	return nums[k-1]
}
