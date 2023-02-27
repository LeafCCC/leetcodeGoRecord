package leetcode

import "sort"

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)

	begin := 0

	for i := (len(nums) + 1) / 2; i < len(nums); i++ {
		if nums[begin]*2 <= nums[i] {
			begin++
		}
	}

	return begin * 2
}
