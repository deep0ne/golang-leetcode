// https://leetcode.com/problems/binary-tree-preorder-traversal/description/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, nodes []int) []int {
	nodes = append(nodes, root.Val)
	if root.Left != nil {
		nodes = helper(root.Left, nodes)
	}
	if root.Right != nil {
		nodes = helper(root.Right, nodes)
	}
	return nodes
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return helper(root, []int{})
}
