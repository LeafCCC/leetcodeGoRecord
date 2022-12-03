package leetcode

func binarySearch(nums []int, begin int, end int, target int) int {
	if begin > end {
		return -1
	}

	mid := begin + (end-begin)/2

	if nums[mid] < target {
		a := binarySearch(nums, mid+1, end, target)
		if a != -1 {
			return a
		}

		if nums[mid] < nums[begin] {
			b := binarySearch(nums, begin, mid-1, target)
			if b != -1 {
				return b
			}
		}
		return -1

	}

	if nums[mid] > target {
		a := binarySearch(nums, begin, mid-1, target)
		if a != -1 {
			return a
		}

		if nums[mid] > nums[end] {
			b := binarySearch(nums, mid+1, end, target)
			if b != -1 {
				return b
			}
		}
		return -1
	}

	return mid
}

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}
