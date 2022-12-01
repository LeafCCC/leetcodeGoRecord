package leetcode

import "math"

func divide(dividend int, divisor int) int {
	//排除溢出的情况
	//结果大于IntMax
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	//除数太大
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}

	//被除数是0
	if dividend == 0 {
		return 0
	}

	//除数与被除数全部变为负数考虑
	sign := true
	if dividend > 0 && divisor > 0 {
		dividend = -dividend
		divisor = -divisor
	} else if dividend < 0 && divisor > 0 {
		sign = false
		divisor = -divisor
	} else if dividend > 0 && divisor < 0 {
		sign = false
		dividend = -dividend
	}

	now := divisor
	rcd := []int{now}
	for now >= dividend-now {
		now = now + now
		rcd = append(rcd, now)
	}

	res := 0

	for i := len(rcd) - 1; i >= 0; i-- {
		if rcd[i] >= dividend {
			res = res | 1<<i
			dividend -= rcd[i]
		}
	}

	if !sign {
		return -res
	}
	return res
}
