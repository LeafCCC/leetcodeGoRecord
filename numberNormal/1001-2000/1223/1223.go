package leetcode

func dieSimulator(n int, rollMax []int) int {
	M := 1000000000 + 7

	//其实这题想出状态方程就不难
	//dp数组记录以某个连续数字结尾的个数 以1为例 rollMax[0] = 2
	//则有dp[0] dp[1] 分别表示以1个1结尾 和 连续两个1结尾
	//则n+1的情况下 dp[1] = dp[0] 而dp[0]等于总和减去除1结尾的情况
	//因此需要一个iSum来记录以某个数字结尾的总和

	//dp[i][j]表示以连续j个数字i+1结尾的个数
	dp := make([][]int, 6)

	for i := range dp {
		dp[i] = make([]int, rollMax[i])
		dp[i][0] = 1
	}

	//iSum记录以i+1结尾的个数和
	iSum := make([]int, 6)
	for i := range iSum {
		iSum[i] = 1
	}

	//res记录从n=1开始的结果
	res := 6

	for i := 2; i <= n; i++ {
		tmpR := res
		res = 0
		for j := range dp {
			for k := len(dp[j]) - 1; k > 0; k-- {
				dp[j][k] = dp[j][k-1]
			}
			dp[j][0] = (tmpR - iSum[j] + M) % M

			iSum[j] = 0
			for _, val := range dp[j] {
				iSum[j] += val
			}
			iSum[j] %= M
			res += iSum[j]
		}
		res %= M
	}

	return res
}
