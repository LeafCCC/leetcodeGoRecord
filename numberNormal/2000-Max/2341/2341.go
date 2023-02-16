package leetcode

func numberOfPairs(nums []int) []int {
	ans := make([]int, 2)
	record := make(map[int]int)

	for _, val := range nums {
		record[val]++
	}

	for _, val := range record {
		ans[0] += val / 2
		ans[1] += val % 2
	}

	return ans
}
