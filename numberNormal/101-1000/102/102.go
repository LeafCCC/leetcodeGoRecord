package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	res := [][]int{}

	for len(q) != 0 {
		tmp := []*TreeNode{}
		r := []int{}
		for _, node := range q {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}

			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
			r = append(r, node.Val)
		}
		res = append(res, r)
		q = tmp
	}
	return res
}
