package leetcode

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	//摆烂看题解了

	//使用一个前缀和W存储前i个箱子的重量
	n := len(boxes)

	W := make([]int, n+1)

	for i := 0; i < n; i++ {
		W[i+1] = W[i] + boxes[i][1]
	}

	//使用一个前缀和F存储到末尾的行程数
	F := make([]int, n)

	for i := 1; i < n; i++ {
		if boxes[i][0] != boxes[i-1][0] {
			F[i] = F[i-1] + 1
		} else {
			F[i] = F[i-1]
		}
	}

	//dp保持前i个箱子的结果
	//dp[i+1] = dp[j] + f(boxes, j, i)
	dp := make([]int, n+1)

	//记dp[j] - F[j] = G[j]
	//dp[i+1] = min(G[j]) + F[i] + 2
	// W[j] >= W[i+1] - maxWeight
	// j >= i - maxBoxes + 1
	//维护一个单调队列保存递增的G
	G := [][]int{}
	//第一位是值 第二位是下标值即j

	for i := 0; i < n; i++ {
		Gi := dp[i] - F[i]
		for len(G) != 0 && G[len(G)-1][0] >= Gi {
			G = G[:len(G)-1]
		}
		G = append(G, []int{Gi, i})

		for true {
			j := G[0][1]
			Gj := G[0][0]
			if W[j] >= W[i+1]-maxWeight && j >= i-maxBoxes+1 {
				dp[i+1] = Gj + F[i] + 2
				break
			}
			G = G[1:]
		}
	}

	//这是复杂度为O(n²)的做法
	// for i := 0; i < n; i++{
	//     begin := 0
	//     if i - maxBoxes + 1 >= 0{
	//         begin = i - maxBoxes + 1
	//     }
	//     for j := begin; j <= i; j++{
	//         if (W[i+1] - W[j]) <= maxWeight{
	//             tmp := dp[j] - F[j] + F[i]  + 2
	//             if dp[i+1] == 0 ||  dp[i+1] > tmp {
	//                 dp[i+1] = tmp
	//             }
	//         }
	//     }
	// }

	return dp[n]
}
