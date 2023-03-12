package leetcode

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}

	l, r := 1, x
	for l+1 < r {
		mid := (l + r) / 2
		if mid*mid > x {
			r = mid
		} else {
			l = mid
		}
	}

	return l
}
