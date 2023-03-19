package leetcode

func numTrees(n int) int {
	var help func(int, int) int

	type p struct {
		s, e int
	}
	record := make(map[p]int)

	help = func(start, end int) int {
		if record[p{start, end}] != 0 {
			return record[p{start, end}]
		}

		if start > end {
			return 1
		}

		res := 0

		for i := start; i <= end; i++ {
			l := help(start, i-1)
			r := help(i+1, end)
			res += l * r
		}

		record[p{start, end}] = res

		return res
	}

	return help(1, n)
}
