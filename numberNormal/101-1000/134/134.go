package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	rcd := make([]int, len(gas))
	sum := 0

	for i := range rcd {
		rcd[i] = gas[i] - cost[i]
		sum += rcd[i]
	}

	if sum < 0 {
		return -1
	}

	sum = rcd[0]
	left := 0
	right := 0

	for !(right+1 == left) && !(left == 0 && right == len(rcd)-1) {
		if sum >= 0 {
			right++
			if right == len(rcd) {
				right = 0
			}
			sum += rcd[right]
		} else {
			if left == right {
				left++
				right++
				sum = rcd[left]
			} else {
				sum -= rcd[left]
				left++
			}
		}
	}

	return left
}
