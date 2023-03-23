package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	// if root == nil{
	//     return [][]int{}
	// }

	stack := []*TreeNode{root}
	res := [][]int{}
	now := true

	for len(stack) != 0 {
		tmp := []*TreeNode{}
		tmpR := []int{}

		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i] == nil {
				continue
			}
			tmpR = append(tmpR, stack[i].Val)
			if now {
				tmp = append(tmp, stack[i].Left)
				tmp = append(tmp, stack[i].Right)
			} else {
				tmp = append(tmp, stack[i].Right)
				tmp = append(tmp, stack[i].Left)
			}
		}

		if len(tmpR) != 0 {
			res = append(res, tmpR)
		}
		stack = tmp
		now = !now
	}

	return res
}
