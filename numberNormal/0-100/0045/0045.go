package leetcode

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	right := nums[0]
	now := 0
	step := 0

	for right < n-1 {
		step++
		tmpNow, tmpR := now, right
		for i := now + 1; i <= right; i++ {
			if i+nums[i] > tmpR {
				tmpNow = i
				tmpR = i + nums[i]
			}
		}
		right = tmpR
		now = tmpNow
	}

	return step + 1
}
