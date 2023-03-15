package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	res := &ListNode{Val: 0, Next: head}
	before := res
	now := head

	for now != nil && now.Next != nil {
		if now.Val == now.Next.Val {
			next := now.Next.Next
			for next != nil && next.Val == now.Val {
				next = next.Next
			}
			before.Next = next
			now = next
		} else {
			before = now
			now = now.Next
		}
	}
	return res.Next
}
