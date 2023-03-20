package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	res := []*TreeNode{}

	var help func(root *TreeNode)

	help = func(root *TreeNode) {
		if root == nil {
			return
		}
		help(root.Left)
		res = append(res, root)
		help(root.Right)
	}

	help(root)

	first, second := &TreeNode{}, &TreeNode{}
	for i := range res {
		if res[i].Val > res[i+1].Val {
			first = res[i]
			second = res[i+1]
			for j := i + 2; j < len(res); j++ {
				if res[j].Val < res[j-1].Val {
					second = res[j]
					break
				}
			}
			break
		}
	}

	first.Val, second.Val = second.Val, first.Val
}
