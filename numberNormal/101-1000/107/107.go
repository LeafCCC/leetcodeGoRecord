package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	order := [][]int{}

	if root == nil {
		return order
	}

	q := []*TreeNode{root}

	for len(q) != 0 {
		tmp := []int{}
		tmpq := []*TreeNode{}

		for _, node := range q {
			if node.Left != nil {
				tmpq = append(tmpq, node.Left)
			}

			if node.Right != nil {
				tmpq = append(tmpq, node.Right)
			}

			tmp = append(tmp, node.Val)
		}

		q = tmpq
		order = append(order, tmp)
	}

	res := [][]int{}
	for i := len(order) - 1; i >= 0; i-- {
		res = append(res, order[i])
	}
	return res
}
