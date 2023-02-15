package leetcode

func balancedString(s string) int {
	record := map[byte]int{}
	target := len(s) / 4

	for _, c := range []byte(s) {
		record[c]++
	}

	check := func(record map[byte]int) bool {
		for _, val := range record {
			if val > target {
				return false
			}
		}

		return true
	}

	if check(record) {
		return 0
	}

	right, res := -1, len(s)

	for left := range s {
		for !check(record) && right < len(s)-1 {
			right++
			record[s[right]]--
		}

		if !check(record) {
			return res
		}

		if right-left+1 < res {
			res = right - left + 1
		}

		record[s[left]]++
	}

	return res
}
