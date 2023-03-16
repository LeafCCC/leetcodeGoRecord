package leetcode

func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}

	gray := grayCode(n - 1)
	res := []int{}
	for i := range gray {
		res = append(res, gray[i])
	}
	for i := range gray {
		res = append(res, gray[len(gray)-i-1]+(1<<(n-1)))
	}
	return res
}
