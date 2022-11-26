package leetcode

import "math"

func myAtoi(s string) int {
	res := 0
	sign := 1 //记录正负号
	con := 0  //记录状态
	for _, i := range s {
		if con == 0 {
			if i == ' ' {
				continue
			} else if i == '+' {
				con += 1
			} else if i == '-' {
				sign = 0
				con += 1
			} else if i >= '0' && i <= '9' {
				res += int(i - '0')
				con += 1
			} else {
				return res
			}
		} else if con == 1 {
			if i >= '0' && i <= '9' {
				res *= 10
				res += int(i - '0')
				if res >= math.MaxInt32+1 {
					if sign == 1 {
						return math.MaxInt32
					} else {
						return math.MinInt32
					}
				}
			} else {
				break
			}
		}
	}

	if res >= math.MaxInt32+1 {
		if sign == 1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}

	if sign == 0 {
		return -res
	}

	return res
}
