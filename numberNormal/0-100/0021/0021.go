package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	now := head

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			now.Next = &ListNode{Val: list2.Val}
			now = now.Next
			list2 = list2.Next
		} else {
			now.Next = &ListNode{Val: list1.Val}
			now = now.Next
			list1 = list1.Next
		}
	}

	if list1 != nil {
		now.Next = list1
	} else if list2 != nil {
		now.Next = list2
	}

	return head.Next
}

// 递归
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists2(list1, list2.Next)
	return list2
}
