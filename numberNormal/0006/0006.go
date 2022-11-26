package leetcode

//找规律
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)

	res := make([]byte, 0, n)

	//第一行
	for i := 0; i < n; i += 2*numRows - 2 {
		res = append(res, s[i])
	}

	//中间行
	for i := 1; i < numRows-1; i++ {
		count := 1
		now := i
		for now < n {
			res = append(res, s[now])

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
		res = append(res, s[i])
	}

	return string(res)
}
