// https://leetcode.com/problems/range-sum-of-bst/description/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum, l, r := 0, 0, 0
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}
	if low <= root.Val {
		l = rangeSumBST(root.Left, low, high)
	}
	if high >= root.Val {
		r = rangeSumBST(root.Right, low, high)
	}
	return sum + l + r
}
