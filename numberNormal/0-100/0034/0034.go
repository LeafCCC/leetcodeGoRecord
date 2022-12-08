package leetcode

func binarySearch(nums *[]int, begin int, end int, target int) int {
	if begin > end {
		return -1
	}

	mid := begin + (end-begin)/2

	if (*nums)[mid] == target {
		return mid
	}

	if (*nums)[mid] > target {
		return binarySearch(nums, begin, mid-1, target)
	}

	return binarySearch(nums, mid+1, end, target)
}

func searchRange(nums []int, target int) []int {
	where := binarySearch(&nums, 0, len(nums)-1, target)

	if where == -1 {
		return []int{-1, -1}
	}

	begin, end := where, where

	for begin >= 1 && nums[begin-1] == target {
		begin--
	}

	for end < len(nums)-1 && nums[end+1] == target {
		end++
	}

	return []int{begin, end}
}
