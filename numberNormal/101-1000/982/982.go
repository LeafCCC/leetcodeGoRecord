package leetcode

func countTriplets(nums []int) int {
	record := make(map[int]int)

	for _, i := range nums {
		for _, j := range nums {
			record[i&j]++
		}
	}

	res := 0
	for _, k := range nums {
		for key, val := range record {
			if k&key == 0 {
				res += val
			}
		}
	}

	return res
}
