package leetcode

//想到了滑动窗口 但是没写好 还是看了官方的
func findSubstring(s string, words []string) (res []int) {
	m := len(words)    //words的数量
	n := len(words[0]) //每个word的长度
	ls := len(s)       //string s的长度

	for i := 0; i < n; i++ {
		if i+m*n > ls {
			break
		}

		differ := map[string]int{}
		for _, k := range words {
			differ[k]++
		}

		for j := i; j < i+m*n; j += n {
			word := s[j : j+n]
			differ[word]--

			if differ[word] == 0 {
				delete(differ, word)
			}
		}

		if len(differ) == 0 {
			res = append(res, i)
		}

		for start := i + n; start+n*m <= ls; start += n {
			word1 := s[start-n : start]
			word2 := s[start+n*m-n : start+n*m]
			differ[word1]++
			differ[word2]--

			if differ[word1] == 0 {
				delete(differ, word1)
			}

			if differ[word2] == 0 {
				delete(differ, word2)
			}

			if len(differ) == 0 {
				res = append(res, start)
			}
		}
	}
	return
}
