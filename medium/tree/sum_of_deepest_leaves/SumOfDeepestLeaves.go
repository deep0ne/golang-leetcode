// https://leetcode.com/problems/deepest-leaves-sum/description/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	result := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		sum := 0
		n := len(nodes)
		for i := 0; i < n; i++ {
			node := nodes[0]
			nodes = nodes[1:]
			sum += node.Val
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		result = sum
	}
	return result
}
