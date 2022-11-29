//注意溢出即可

package leetcode

import "math"

func reverse(x int) int {
	if x == math.MinInt32 {
		return 0
	}
	sign := 1
	if x < 0 {
		sign = 0
		x = -x
	}

	res := 0
	for x > 0 && res < math.MaxInt32/10 {
		res *= 10
		res += x % 10
		x = x / 10
	}
	if x != 0 {
		if res == math.MaxInt32/10 {
			if x > 7 {
				return 0
			} else {
				res *= 10
				res += x
			}
		} else {
			return 0
		}
	}
	if sign == 1 {
		return res
	}
	return -res

}
