package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{Next: head}
	now := tmp
	sz := 0

	for now != nil {
		now = now.Next
		sz++
	}
	position := sz - n - 1

	now = tmp
	for position > 0 {
		now = now.Next
		position--
	}

	now.Next = now.Next.Next

	return tmp.Next
}
