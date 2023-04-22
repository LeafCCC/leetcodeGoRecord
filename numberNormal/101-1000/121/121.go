package leetcode

func maxProfit(prices []int) int {
	minP := prices[0]
	res := 0

	for i := 1; i < len(prices); i++ {
		res = max(prices[i]-minP, res)
		minP = min(minP, prices[i])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
