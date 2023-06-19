// https://leetcode.com/problems/validate-binary-search-tree/description/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var tmp *TreeNode
	nodesStack := make([]*TreeNode, 0)
	for len(nodesStack) > 0 || root != nil {
		for root != nil {
			nodesStack = append(nodesStack, root)
			root = root.Left
		}
		root = nodesStack[len(nodesStack)-1]
		nodesStack = nodesStack[:len(nodesStack)-1]
		if tmp != nil && tmp.Val >= root.Val {
			return false
		}
		tmp = root
		root = root.Right
	}
	return true
}
