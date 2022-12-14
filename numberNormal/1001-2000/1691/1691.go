package leetcode

import "sort"

func maxHeight(cuboids [][]int) int {
	for i := range cuboids {
		sort.Ints(cuboids[i])
	}

	//第一位宽 第二位长 第三位高
	//按宽的大小降序排序
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] == cuboids[j][0] && cuboids[i][1] == cuboids[j][1] {
			return cuboids[i][2] > cuboids[j][2]
		} else if cuboids[i][0] == cuboids[j][0] {
			return cuboids[i][1] > cuboids[j][1]
		}
		return cuboids[i][0] > cuboids[j][0]
	})

	n := len(cuboids)

	//dp[i]记录以第i个纸箱放在最上面的最高高度
	dp := make([]int, n)
	dp[0] = cuboids[0][2]

	for i := 1; i < n; i++ {
		dp[i] = cuboids[i][2]
		for j := 0; j < i; j++ {
			if cuboids[i][2] <= cuboids[j][2] && cuboids[i][1] <= cuboids[j][1] {
				dp[i] = max(dp[i], dp[j]+cuboids[i][2])
			}
		}
	}

	res := dp[0]

	for _, val := range dp {
		res = max(res, val)
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
