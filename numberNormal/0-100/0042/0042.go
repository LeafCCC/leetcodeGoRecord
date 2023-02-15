package leetcode

func trap(height []int) int {
	res := 0

	begin := 0
	tmp := 0

	for i := 1; i < len(height); i++ {
		if height[i] < height[begin] {
			tmp += height[begin] - height[i]
		} else {
			begin = i
			res += tmp
			tmp = 0
		}
	}

	end := len(height) - 1
	tmp = 0

	for i := len(height) - 2; i >= begin; i-- {
		if height[i] < height[end] {
			tmp += height[end] - height[i]
		} else {
			end = i
			res += tmp
			tmp = 0
		}
	}

	return res
}
