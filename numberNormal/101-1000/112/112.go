package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func help(root *TreeNode, nowSum int, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		return nowSum == targetSum
	}

	if root.Left == nil {
		return help(root.Right, nowSum+root.Right.Val, targetSum)
	}

	if root.Right == nil {
		return help(root.Left, nowSum+root.Left.Val, targetSum)
	}

	return help(root.Left, nowSum+root.Left.Val, targetSum) || help(root.Right, nowSum+root.Right.Val, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return help(root, root.Val, targetSum)
}
