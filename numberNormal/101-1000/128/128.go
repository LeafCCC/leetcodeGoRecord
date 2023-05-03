package leetcode

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	record := make(map[int]bool)
	for _, val := range nums {
		record[val] = true
	}

	res := 1
	for _, val := range nums {
		if record[val-1] {
			continue
		}

		tmp := 1
		now := val + 1
		for record[now] {
			now++
			tmp++
		}
		res = max(res, tmp)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
