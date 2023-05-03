package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func help(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}

	return help(root.Left, sum*10+root.Val) + help(root.Right, sum*10+root.Val)
}

func sumNumbers(root *TreeNode) int {
	return help(root, 0)
}
