package leetcode

func romanToInt(s string) (res int) {
	rcd := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	t := []byte(s)
	for i := 0; i < len(t)-1; i++ {
		if rcd[t[i]] < rcd[t[i+1]] {
			res -= rcd[t[i]]
		} else {
			res += rcd[t[i]]
		}
	}
	res += rcd[t[len(t)-1]]
	return
}
