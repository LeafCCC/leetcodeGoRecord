package main

func beautySum(s string) (res int) {
	n := len(s)

	for i := 0; i < n; i++ {
		rcd := make(map[byte]int)
		rcd[s[i]]++
		for j := i + 1; j < n; j++ {
			rcd[s[j]]++
			min, max := rcd[s[j]], rcd[s[j]]
			for _, val := range rcd {
				if val < min {
					min = val
				}
				if val > max {
					max = val
				}
			}
			res += max - min
		}

	}
	return
}
