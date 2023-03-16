package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	smaller, bigger := []int{}, []int{}

	now := head
	for now != nil {
		if now.Val < x {
			smaller = append(smaller, now.Val)
		} else {
			bigger = append(bigger, now.Val)
		}

		now = now.Next
	}

	res := &ListNode{0, nil}
	now = res

	for _, val := range smaller {
		now.Next = &ListNode{val, nil}
		now = now.Next
	}

	for _, val := range bigger {
		now.Next = &ListNode{val, nil}
		now = now.Next
	}
	return res.Next
}
