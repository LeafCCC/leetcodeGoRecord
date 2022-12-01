package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseK(head *ListNode, k int) *ListNode {
	lists := []*ListNode{}
	now := head.Next

	for i := 0; i < k; i++ {
		if now == nil {
			return nil
		}
		lists = append(lists, now)
		now = now.Next
	}

	head.Next = lists[k-1]
	lists[0].Next = lists[k-1].Next

	for i := 1; i < k; i++ {
		lists[i].Next = lists[i-1]
	}

	return lists[0]
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	res := &ListNode{Next: head}
	tmp := res

	for tmp != nil {
		tmp = reverseK(tmp, k)
	}

	return res.Next
}
