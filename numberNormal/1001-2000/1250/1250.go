package leetcode

func isGoodArray(nums []int) bool {
	var gcd func(a, b int) int

	gcd = func(a, b int) int {
		if a%b == 0 {
			return b
		}
		return gcd(b, a%b)
	}

	t := nums[0]
	for _, val := range nums {
		t = gcd(val, t)
		if t == 1 {
			return true
		}
	}

	return false
}
