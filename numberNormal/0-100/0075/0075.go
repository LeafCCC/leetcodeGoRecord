package leetcode

func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for l < len(nums) && nums[l] == 0 {
		l++
	}
	for r >= 0 && nums[r] == 2 {
		r--
	}

	for i := l; i <= r; i++ {
		for nums[i] != 1 && i >= l && i <= r {
			if nums[i] == 0 {
				nums[i], nums[l] = nums[l], nums[i]
				l++
			} else {
				nums[i], nums[r] = nums[r], nums[i]
				r--
			}
		}
	}

}
