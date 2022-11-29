package leetcode

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxArea(height []int) (res int) {
	begin, end := 0, len(height)-1
	tmp := 0

	for begin < end {
		if height[begin] < height[end] {
			tmp = (end - begin) * height[begin]
			begin++
		} else {
			tmp = (end - begin) * height[end]
			end--
		}

		res = max(res, tmp)
	}

	return
}
