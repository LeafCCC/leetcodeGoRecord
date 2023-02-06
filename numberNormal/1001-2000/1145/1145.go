package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var caculate func(root *TreeNode) int
	caculate = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + caculate(root.Left) + caculate(root.Right)
	}

	var find func(root *TreeNode, x int) *TreeNode
	find = func(root *TreeNode, x int) *TreeNode {

		if root.Val == x {
			return root
		}
		if root.Left != nil {
			t := find(root.Left, x)
			if t != nil {
				return t
			}
		}
		if root.Right != nil {
			t := find(root.Right, x)
			if t != nil {
				return t
			}
		}
		return nil
	}

	if x == root.Val {
		return abs(caculate(root.Left)-caculate(root.Right)) > 1
	}

	red := find(root, x)
	number1 := caculate(red)
	number2 := max(caculate(red.Left), caculate(red.Right))
	return n-number1 > number1 || number2 > n-number2
}
