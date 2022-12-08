//普通的二分

package leetcode

func searchInsert(nums []int, target int) int {
	begin, end := 0, len(nums)-1
	if nums[begin] >= target {
		return 0
	}

	if nums[end] <= target {
		if nums[end] == target {
			return end
		}
		return end + 1
	}

	for begin+1 < end {
		mid := begin + (end-begin)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			end = mid
		} else {
			begin = mid
		}
	}
	return end
}
