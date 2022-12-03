package leetcode

func secondHighest(s string) int {
	res := []int{}

	for _, i := range s {
		if i >= '0' && i <= '9' {
			tmp := int(i - '0')

			if len(res) == 0 {
				res = append(res, tmp)
			} else if len(res) == 1 && res[0] != tmp {
				res = append(res, tmp)
				if tmp < res[0] {
					res[0], res[1] = res[1], res[0]
				}
			} else if len(res) == 2 {
				if tmp > res[0] && tmp < res[1] {
					res[0] = tmp
				} else if tmp > res[1] {
					res[0] = res[1]
					res[1] = tmp
				}
			}
		}
	}

	if len(res) < 2 {
		return -1
	}
	return res[0]
}
