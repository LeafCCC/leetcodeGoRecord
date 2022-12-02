package leetcode

import "sort"

func nextPermutation(nums []int) {
	//找规律 逆序遍历 直到找到更小的值
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := n - 1; j > i; j-- {
				if nums[j] > nums[i] {
					nums[j], nums[i] = nums[i], nums[j]
					sort.Ints(nums[i+1:])
					return
				}
			}
		}
	}
	sort.Ints(nums)
	return
}
