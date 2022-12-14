package leetcode

func checkIfPangram(sentence string) bool {
	rcd := make(map[byte]bool)

	for i := 0; i < 26; i++ {
		rcd[byte('a'+i)] = false
	}

	for i := range sentence {
		rcd[sentence[i]] = true
	}

	for i := range rcd {
		if !rcd[i] {
			return false
		}
	}

	return true
}
