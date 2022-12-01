package leetcode

func removeDuplicates(nums []int) int {
	now := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[now] = nums[i]
			now++
		}
	}

	return now
}
