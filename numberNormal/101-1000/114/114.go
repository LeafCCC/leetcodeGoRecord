package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	var last *TreeNode
	var help func(root *TreeNode)

	help = func(root *TreeNode) {
		if root == nil {
			return
		}

		help(root.Right)
		help(root.Left)

		root.Left = nil
		root.Right = last

		last = root
	}

	help(root)
}
