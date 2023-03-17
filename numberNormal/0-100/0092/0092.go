package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	res := &ListNode{0, head}
	now := res
	for i := 1; i < left; i++ {
		now = now.Next
	}
	begin := now

	now = now.Next
	tmp := []*ListNode{}
	for i := left; i <= right; i++ {
		tmp = append(tmp, now)
		now = now.Next
	}

	begin.Next = tmp[len(tmp)-1]
	for i := len(tmp) - 1; i > 0; i-- {
		tmp[i].Next = tmp[i-1]
	}
	tmp[0].Next = now

	return res.Next
}
