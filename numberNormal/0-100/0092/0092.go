package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 头插法 复杂度O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}

	pre := dummyNode
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	next := cur.Next

	for i := left; i < right; i++ {
		pre.Next, cur.Next, next.Next = next, next.Next, pre.Next
		next = cur.Next
	}

	return dummyNode.Next
}
