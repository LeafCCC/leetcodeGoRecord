package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	tmp := strings.Split(s, " ")

	return len(tmp[len(tmp)-1])
}
