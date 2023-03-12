package leetcode

func plusOne(digits []int) []int {
	count := 1

	for i := len(digits) - 1; i >= 0 && count == 1; i-- {
		val := digits[i] + count
		digits[i] = val % 10
		count = val / 10
	}

	if count == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
