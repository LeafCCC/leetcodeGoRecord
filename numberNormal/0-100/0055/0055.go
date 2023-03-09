package leetcode

func canJump(nums []int) bool {
	now := 0
	n := len(nums)
	steps := nums[0]

	if n == 1 {
		return true
	}

	for steps != 0 {
		if now+steps >= n-1 {
			return true
		}

		max := now + steps
		tmp := now
		for i := 1; i <= steps; i++ {
			if now+i+nums[now+i] >= max {
				max = now + i + nums[now+i]
				tmp = now + i
			}
		}
		now = tmp
		steps = nums[tmp]
	}

	return false
}
