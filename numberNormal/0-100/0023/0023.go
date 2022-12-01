package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个有序链表 即第21题
func mergeTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwo(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwo(l1, l2.Next)
	return l2
}

// 最朴素的解法，即从第一个合并到最后一个
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	res := lists[0]

	for i := 1; i < len(lists); i++ {
		res = mergeTwo(res, lists[i])
	}

	return res
}

// 分治法
func f(lists []*ListNode, begin int, end int) *ListNode {
	if begin == end {
		return lists[begin]
	}

	if begin+1 == end {
		return mergeTwo(lists[begin], lists[end])
	}

	mid := (begin + end) / 2

	a := f(lists, begin, mid)
	b := f(lists, mid+1, end)
	return mergeTwo(a, b)
}

func mergeKLists2(lists []*ListNode) *ListNode {
	end := len(lists) - 1

	if end == -1 {
		return nil
	}
	return f(lists, 0, end)
}
