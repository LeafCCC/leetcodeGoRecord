package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	max := math.MaxInt64
	min := math.MinInt64

	var help func(root *TreeNode, lower, bigger int) bool

	help = func(root *TreeNode, lower, bigger int) bool {
		if root == nil {
			return true
		}

		if root.Val >= bigger {
			return false
		}

		if root.Val <= lower {
			return false
		}

		return help(root.Left, lower, root.Val) && help(root.Right, root.Val, bigger)
	}

	return help(root, min, max)
}
