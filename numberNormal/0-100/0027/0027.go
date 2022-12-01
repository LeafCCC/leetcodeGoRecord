package leetcode

func removeElement(nums []int, val int) int {
	now := 0

	for i := range nums {
		if nums[i] != val {
			nums[now] = nums[i]
			now++
		}
	}

	return now
}
