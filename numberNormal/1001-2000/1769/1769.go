package leetcode

func minOperations(boxes string) (res []int) {
	left, right := 0, 0  //记录左右个数
	lStep, rStep := 0, 0 //记录左右操作数
	n := len(boxes)

	//初始时统计下标为0时右边的个数与操作数
	for i := 1; i < n; i++ {
		if boxes[i] == '1' {
			rStep += i
			right++
		}
	}

	res = append(res, rStep)

	//从下标为1开始遍历
	for i := 1; i < n; i++ {
		if boxes[i-1] == '1' {
			left++
		}
		lStep += left
		rStep -= right
		if boxes[i] == '1' {
			right--
		}
		res = append(res, lStep+rStep)
	}

	return
}
