package leetcode

func alphabetBoardPath(target string) string {
	res := []byte{}

	x, y := 0, 0
	for _, c := range target {
		cx, cy := int(c-'a')%5, int(c-'a')/5

		if y > cy {
			for i := 0; i < y-cy; i++ {
				res = append(res, 'U')
			}
		}

		if x > cx {
			for i := 0; i < x-cx; i++ {
				res = append(res, 'L')
			}
		} else {
			for i := 0; i < cx-x; i++ {
				res = append(res, 'R')
			}
		}

		if y < cy {
			for i := 0; i < cy-y; i++ {
				res = append(res, 'D')
			}
		}
		res = append(res, '!')
		x, y = cx, cy
	}

	return string(res)
}
