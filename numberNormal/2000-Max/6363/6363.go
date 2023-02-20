package leetcode

func minOperations(n int) int {
	step := 0

	for n != 0 {
		if n&1 == 1 {
			if (n>>1)&1 != 1 {
				step++
				n = n >> 1
			} else {
				n = n + 1
				step++
			}
		} else {
			n = n >> 1
		}
	}

	return step
}
