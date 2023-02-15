package leetcode

func longestWPI(hours []int) int {
	s := make([]int, len(hours)+1)
	stack := []int{0}

	for i := range hours {
		if hours[i] > 8 {
			s[i+1] = s[i] + 1
		} else {
			s[i+1] = s[i] - 1
		}

		if s[i+1] < s[stack[len(stack)-1]] {
			stack = append(stack, i+1)
		}
	}

	res := 0
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	for i := len(s) - 1; i > 0; i-- {
		for len(stack) > 0 && s[i] > s[stack[len(stack)-1]] {
			res = max(i-stack[len(stack)-1], res)
			stack = stack[:len(stack)-1]
		}
	}

	return res
}
