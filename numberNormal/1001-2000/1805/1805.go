package leetcode

func numDifferentIntegers(word string) (res int) {
	rcd := map[string]bool{}
	begin := 0
	now := false
	if word[0] >= '0' && word[0] <= '9' {
		now = true
	}

	for i := 1; i < len(word); i++ {
		if word[i] >= '0' && word[i] <= '9' {
			if !now {
				begin = i
				now = true
			}
		} else if now {
			tmp := word[begin:i]
			for len(tmp) > 1 && tmp[0] == '0' {
				tmp = tmp[1:]
			}
			if _, isPresent := rcd[tmp]; !isPresent {
				rcd[tmp] = true
				res++
			}
			now = false
		}
	}

	if word[len(word)-1] >= '0' && word[len(word)-1] <= '9' {
		tmp := word[begin:len(word)]
		for len(tmp) > 1 && tmp[0] == '0' {
			tmp = tmp[1:]
		}
		if _, isPresent := rcd[tmp]; !isPresent {
			res++
		}
	}

	return
}
