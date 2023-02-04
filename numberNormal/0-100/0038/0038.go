package leetcode

import "strconv"

func countAndSay(n int) string {
	f := func(s string) string {
		k := 1
		res := ""
		for i := 1; i < len(s); i++ {
			if s[i] == s[i-1] {
				k++
			} else {
				res += strconv.Itoa(k) + s[i-1:i]
				k = 1
			}
		}
		res += strconv.Itoa(k) + s[len(s)-1:len(s)]
		return res
	}

	ans := "1"

	for i := 1; i < n; i++ {
		ans = f(ans)
	}

	return ans
}
