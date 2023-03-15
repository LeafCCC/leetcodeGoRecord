package leetcode

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 1 {
		return nums[0] == target
	} else if n == 2 {
		return nums[0] == target || nums[1] == target
	}

	begin, end := 0, n-1
	mid := begin + (end-begin)/2

	if target == nums[mid] {
		return true
	}

	if nums[0] < nums[mid] {
		if target < nums[mid] {
			if nums[mid] <= nums[end] {
				return search(nums[:mid], target)
			} else {
				if target >= nums[begin] {
					return search(nums[:mid], target)
				} else {
					return search(nums[mid+1:], target)
				}
			}

		}
		if target > nums[mid] {
			return search(nums[mid+1:], target)
		}
	}

	if nums[0] > nums[mid] {
		if target >= nums[0] || target <= nums[mid] {
			return search(nums[:mid+1], target)
		} else {
			return search(nums[mid+1:], target)
		}
	}

	if nums[0] == nums[mid] {
		return search(nums[:mid], target) || search(nums[mid+1:], target)
	}

	return false
}
