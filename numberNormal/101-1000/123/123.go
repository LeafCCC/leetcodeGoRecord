package leetcode

func maxProfit(prices []int) int {
	//用两个dp数组记录 分别表示从左到右的最大一笔交易 和 从右到左的最大一笔交易
	n := len(prices)
	dp1, dp2 := make([]int, n), make([]int, n)

	pmin := prices[0]
	for i := 1; i < n; i++ {
		if prices[i] < pmin {
			pmin = prices[i]
		}

		dp1[i] = max(dp1[i-1], prices[i]-pmin)
	}

	pmax := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		if prices[i] > pmax {
			pmax = prices[i]
		}

		dp2[i] = max(dp2[i+1], pmax-prices[i])
	}

	res := dp1[n-1]
	for i := 0; i < n-1; i++ {
		res = max(res, dp1[i]+dp2[i+1])
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
