package leetcode

import "sort"

func f(nums1 *[]int, nums2 *[]int, count1 int, count2 int) (res int) {
	n1, n2 := len(*nums1), len(*nums2)
	p1, p2 := 0, n2-1
	for count1 < count2 {
		if p1 == n1 && p2 == -1 {
			res = -1
			return
		}

		if p1 == n1 || (p2 != -1 && 6-(*nums1)[p1] <= (*nums2)[p2]-1) {
			count2 -= (*nums2)[p2] - 1
			res++
			p2--
		} else if p2 == -1 || 6-(*nums1)[p1] > (*nums2)[p2]-1 {
			count1 += 6 - (*nums1)[p1]
			res++
			p1++
		}
	}
	return
}

func minOperations(nums1 []int, nums2 []int) int {
	count1 := 0
	count2 := 0

	sort.Ints(nums1)
	sort.Ints(nums2)

	for _, i := range nums1 {
		count1 += i
	}

	for _, i := range nums2 {
		count2 += i
	}

	if count1 == count2 {
		return 0
	}

	if count1 < count2 {
		return f(&nums1, &nums2, count1, count2)
	}

	return f(&nums2, &nums1, count2, count1)
}
