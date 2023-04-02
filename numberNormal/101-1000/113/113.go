package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, targetSum int) [][]int {

	res := [][]int{}
	tmp := []int{}

	var help func(*TreeNode, int, int)
	help = func(root *TreeNode, targetSum, now int) {
		if root == nil {
			return
		}

		now += root.Val
		tmp = append(tmp, root.Val)
		defer func() {
			tmp = tmp[:len(tmp)-1]
		}()

		if now == targetSum && root.Left == nil && root.Right == nil {
			res = append(res, append([]int{}, tmp...))
			return
		}

		help(root.Left, targetSum, now)
		help(root.Right, targetSum, now)
	}

	help(root, targetSum, 0)
	return res
}
