package leetcode

func minElements(nums []int, limit int, goal int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}

	if sum == goal {
		return 0
	}

	tmp := sum - goal
	if tmp < 0 {
		tmp = -tmp
	}

	if tmp%limit == 0 {
		return tmp / limit
	}

	return tmp/limit + 1
}
