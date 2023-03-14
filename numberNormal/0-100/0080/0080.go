package leetcode

func removeDuplicates(nums []int) int {
	now := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[now] = nums[i]
			now++
			count = 1
		} else if count != 2 {
			count++
			nums[now] = nums[i]
			now++
		}
	}

	return now
}
