// https://leetcode.com/problems/binary-tree-postorder-traversal/description/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, nodes []int) []int {
	if root.Left != nil {
		nodes = helper(root.Left, nodes)
	}
	if root.Right != nil {
		nodes = helper(root.Right, nodes)
	}
	nodes = append(nodes, root.Val)
	return nodes
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return helper(root, []int{})
}
