package leetcode

import "strconv"

func getLucky(s string, k int) int {
	res := ""

	for i := range s {
		res += strconv.Itoa(int(s[i]-'a') + 1)
	}

	for i := 1; i <= k; i++ {
		cnt := 0
		for j := range res {
			cnt += int(res[j] - '0')
		}
		res = strconv.Itoa(cnt)
	}

	final, _ := strconv.Atoi(res)

	return final
}
