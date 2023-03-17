package leetcode

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) (res []string) {
	tmp := []string{}
	n := len(s)

	var backTrack func(now int)

	backTrack = func(now int) {
		if len(tmp) > 4 {
			return
		}

		if now == n {
			if len(tmp) == 4 {
				t := strings.Join(tmp, ".")
				res = append(res, t)
			}
			return
		}

		if len(tmp) < 4 {
			tmp = append(tmp, s[now:now+1])
			backTrack(now + 1)
			tmp = tmp[:len(tmp)-1]
		}

		if len(tmp) != 0 && tmp[len(tmp)-1] != "0" {
			before := tmp[len(tmp)-1]
			tmp[len(tmp)-1] += s[now : now+1]
			val, _ := strconv.Atoi(tmp[len(tmp)-1])
			if val > 255 {
				return
			}
			backTrack(now + 1)
			tmp[len(tmp)-1] = before
		}

	}

	backTrack(0)

	return res
}
