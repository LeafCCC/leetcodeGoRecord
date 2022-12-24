package leetcode

func minMoves(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	//统计两个1中0的个数
	rcd := []int{}

	t := 0
	for t < len(nums) && nums[t] != 1 {
		t++
	}

	cnt := 0
	for i := t + 1; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt++
		} else {
			rcd = append(rcd, cnt)
			cnt = 0
		}
	}

	//统计前缀和
	tmp := 0
	count := []int{}
	for _, val := range rcd {
		tmp += val
		count = append(count, tmp)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	//先计算第一个窗口的数值
	before := 0

	for i := 0; i < k-1; i++ {
		before += rcd[i] * min(i+1, k-i-1)
	}

	res := before
	begin := 1
	end := 1 + k - 2

	//开始滑动窗口
	for end < len(rcd) {
		left := count[begin+k/2-2] - count[begin-1] + rcd[begin-1]
		right := count[begin-1+k-1] - count[begin-1+(k+1)/2-1]
		before = before - left + right
		res = min(before, res)
		begin++
		end++
	}

	return res

}
