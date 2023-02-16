package leetcode

func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)
	record := []int{}

	for i := range nums {
		if nums[i] < len(nums) && nums[i] > 0 {
			if nums[nums[i]] > 0 && nums[nums[i]] < n {
				record = append(record, nums[nums[i]])
			}
			nums[nums[i]] = nums[i]
		}
	}

	for _, i := range record {
		nums[i] = i
	}

	for i := 1; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}

	return n
}
