package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	end := head
	length := 1

	for end.Next != nil {
		end = end.Next
		length++
	}

	turns := length - (k % length)
	if turns == length {
		return head
	}

	now := head
	for i := 1; i < turns; i++ {
		now = now.Next
	}

	res := now.Next
	end.Next = head
	now.Next = nil

	return res
}
