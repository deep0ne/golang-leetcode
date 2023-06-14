// https://leetcode.com/problems/binary-tree-inorder-traversal/description/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, res []int) []int {
	if root.Left != nil {
		res = helper(root.Left, res)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = helper(root.Right, res)
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root != nil {
		res = helper(root, res)
	}
	return res
}
