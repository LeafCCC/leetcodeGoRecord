package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var recursive func(*TreeNode)
	recursive = func(root *TreeNode) {
		if root == nil {
			return
		}

		recursive(root.Left)
		res = append(res, root.Val)
		recursive(root.Right)
	}

	recursive(root)

	return res
}
