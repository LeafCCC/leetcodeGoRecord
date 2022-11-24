package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	count := 0
	sum := 0
	res := new(ListNode)
	now := res
	for l1 != nil || l2 != nil {
		now.Next = new(ListNode)
		now = now.Next
		sum = count
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum < 10 {
			count = 0
			now.Val = sum
		} else {
			count = 1
			now.Val = sum - 10
		}
	}
	if count == 1 {
		now.Next = &ListNode{Val: 1}
	}
	return res.Next
}
