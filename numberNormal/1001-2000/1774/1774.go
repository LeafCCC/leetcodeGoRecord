package leetcode

import "sort"

//想到了暴力肯定可解 因为这题数据量不超过10
//但是这样写太没意思了 还是照着题解的动态规划解决一下
func closestCost(baseCosts []int, toppingCosts []int, target int) (res int) {
	sort.Ints(baseCosts)
	minBase := baseCosts[0] //寻找基料中的最小值

	if minBase >= target {
		return minBase
	}

	res = 2*target - minBase
	//这个是最大的结果 如果大于这个结果 那么minBase便是更优解

	cost := make([]bool, target+1)

	//先遍历基料
	for _, i := range baseCosts {
		if i <= target {
			cost[i] = true
		} else {
			res = min(res, i)
		}
	}

	//再遍历辅料
	//为什么要逆序遍历？因为前面的结果会影响后面
	//以i=1的辅料为例子
	//遍历前cost[1] = true cost[2] = false cost[3] = false
	//遍历到cost[1]时 cost[1+i] = cost[1] = true
	//但是到cost[2]时 会导致本应该flase的cost[3]变为true
	for _, i := range toppingCosts {
		for j := 0; j < 2; j++ {
			for k := target; k >= 0; k-- {
				//对于大于target的情况 不可能加更多辅料了
				//不然只会增大和target差距
				if cost[k] && k+i > target {
					res = min(res, k+i)
				}

				//对于大于0的情况 说明cost[k]可以由cost[k-i]加上本次辅料获得
				if k-i > 0 {
					cost[k] = cost[k-i] || cost[k]
				}
			}
		}
	}

	//最后逆序遍历dp的数组 注意此时res都是大于target的
	for i := target; i >= 0; i-- {
		if cost[i] && target-i <= res-target {
			return i
		}
	}

	return
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
