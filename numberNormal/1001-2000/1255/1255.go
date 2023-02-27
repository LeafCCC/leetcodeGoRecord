package leetcode

func maxScoreWords(words []string, letters []byte, score []int) int {
	n := len(words)
	letterR := make([]int, 26)
	res := 0

	for _, letter := range letters {
		letterR[int(letter)-int('a')]++
	}

	for s := 1; s < (1 << n); s++ {
		sRecord := make([]int, 26)

		for i := 0; i < n; i++ {
			if s&(1<<i) == 0 {
				continue
			}

			for _, w := range words[i] {
				sRecord[int(w)-int('a')]++
			}
		}

		now := 0
		pass := true

		for i := 0; i < 26; i++ {
			if letterR[i] < sRecord[i] {
				pass = false
				break
			}

			now += score[i] * sRecord[i]
		}

		if pass && now > res {
			res = now
		}
	}

	return res
}
