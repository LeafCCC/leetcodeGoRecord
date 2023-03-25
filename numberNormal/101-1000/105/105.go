package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	record := map[int]int{} //inorder的hash映射
	for i := range inorder {
		record[inorder[i]] = i
	}

	var help func([]int, int, int) *TreeNode

	help = func(preorder []int, start, end int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}

		root := &TreeNode{preorder[0], nil, nil}

		where := record[preorder[0]]
		root.Left = help(preorder[1:where-start+1], start, where-1)
		root.Right = help(preorder[where-start+1:], where+1, end)
		return root
	}

	return help(preorder, 0, len(inorder)-1)
}
