package leetcode

func canChoose(groups [][]int, nums []int) bool {
	gi := 0
	ni := 0

	for ni < len(nums) && gi < len(groups) {
		if groups[gi][0] != nums[ni] {
			ni++
			continue
		}

		i := 1

		for ni+i < len(nums) && i < len(groups[gi]) {
			if nums[ni+i] == groups[gi][i] {
				i++
				continue
			} else {
				break
			}
		}

		if i == len(groups[gi]) {
			ni += i
			gi++
		} else {
			ni++
		}
	}

	return gi >= len(groups)
}
