// https://leetcode.com/problems/validate-binary-search-tree/description/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val || max != nil && root.Val >= max.Val {
		return false
	}
	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}
