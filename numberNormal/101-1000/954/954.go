package leetcode

import "sort"

func canReorderDoubled(arr []int) bool {
	vals := []int{}
	record := make(map[int]int)

	for _, v := range arr {
		if record[v] == 0 {
			vals = append(vals, v)
		}
		record[v]++
	}

	if record[0]%2 != 0 {
		return false
	}
	record[0] = 0

	//按绝对值排序！！
	sort.Slice(vals, func(i, j int) bool {
		return abs(vals[i]) < abs(vals[j])
	})

	for _, v := range vals {
		if record[v] == 0 {
			continue
		}

		if record[v] > record[2*v] {
			return false
		}
		record[2*v] -= record[v]
		record[v] = 0
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
