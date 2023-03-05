package leetcode

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	res := make([]int, len(customers)+1)

	cache := 0

	for i := range customers {
		cache += customers[i]
		if cache == 0 {
			res[i+1] = res[i] - runningCost
		} else if cache <= 4 {
			res[i+1] = res[i] + cache*boardingCost - runningCost
			cache = 0
		} else {
			cache -= 4
			res[i+1] = res[i] + 4*boardingCost - runningCost
		}
	}

	for cache > 0 {
		if cache > 4 {
			res = append(res, res[len(res)-1]+4*boardingCost-runningCost)
			cache -= 4
		} else {
			res = append(res, res[len(res)-1]+cache*boardingCost-runningCost)
			break
		}

	}

	turns := -1
	income := 0

	for i := range res {
		if res[i] > income {
			turns = i
			income = res[i]
		}
	}

	return turns
}
