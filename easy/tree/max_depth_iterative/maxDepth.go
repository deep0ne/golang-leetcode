// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		n := len(nodes)
		for i := 0; i < n; i++ {
			node := nodes[0]
			nodes = nodes[1:]
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		level++
	}
	return level
}
