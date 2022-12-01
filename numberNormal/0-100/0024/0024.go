package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	rcd := &ListNode{Next: head}

	now := rcd

	for now.Next != nil && now.Next.Next != nil {
		a := now.Next
		b := now.Next.Next

		a.Next = b.Next
		b.Next = a
		now.Next = b

		now = a
	}

	return rcd.Next
}
