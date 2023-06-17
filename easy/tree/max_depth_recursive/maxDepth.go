// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
