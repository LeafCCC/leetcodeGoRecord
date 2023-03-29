package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	record := make(map[int]int)
	for i := range inorder {
		record[inorder[i]] = i
	}

	var help func(left, right int, postorder []int) *TreeNode

	//left right 标记inorder的开始和结束
	help = func(left, right int, postorder []int) *TreeNode {
		if len(postorder) == 0 {
			return nil
		}

		val := postorder[len(postorder)-1]
		where := record[val]

		root := &TreeNode{val, nil, nil}
		root.Left = help(left, where-1, postorder[:where-left])
		root.Right = help(where+1, right, postorder[where-left:len(postorder)-1])

		return root
	}

	return help(0, len(inorder)-1, postorder)
}
