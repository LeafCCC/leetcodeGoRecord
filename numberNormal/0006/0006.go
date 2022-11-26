package leetcode

//找规律
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := ""

	n := len(s)

	//第一行
	for i := 0; i < n; i += 2*numRows - 2 {
		res += s[i : i+1]
	}

	//中间行
	for i := 1; i < numRows-1; i++ {
		count := 1
		now := i
		for now < n {
			res += s[now : now+1]

			if count%2 == 1 {
				now += 2*numRows - 2 - 2*i
			} else {
				now += 2 * i
			}
			count++
		}
	}

	//最后一行
	for i := numRows - 1; i < n; i += 2*numRows - 2 {
		res += s[i : i+1]
	}

	return res
}
