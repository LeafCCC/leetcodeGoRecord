package leetcode

func minOperations(nums []int) (res int) {
	before := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] <= before {
			res += before - nums[i] + 1
			before++
		} else {
			before = nums[i]
		}
	}
	return
}
