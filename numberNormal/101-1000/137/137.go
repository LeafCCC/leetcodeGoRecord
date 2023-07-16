func singleNumber(nums []int) int {
	res := make([]int, 32)

	for _, val := range nums {
		for i := 0; i <= 31; i++ {
			if val&(1<<i) != 0 {
				res[i]++
				res[i] %= 3
			}
		}
	}

	target := 0

	for i := range res {
		if i == 31 && res[i] != 0 {
			return target - 1<<31
		} else if res[i] != 0 {
			target += (1 << i)
		}
	}

	return target
}