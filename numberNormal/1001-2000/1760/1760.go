package leetcode

import "sort"

func minimumSize(nums []int, maxOperations int) int {
	sort.Ints(nums)
	left, right := 1, nums[len(nums)-1]

	for left < right {
		t := left + (right-left)/2

		myOps := 0

		for _, val := range nums {
			myOps += (val - 1) / t
		}

		if myOps > maxOperations {
			left = t + 1
		} else {
			right = t
		}
	}

	return right
}
