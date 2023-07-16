package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return true
}
