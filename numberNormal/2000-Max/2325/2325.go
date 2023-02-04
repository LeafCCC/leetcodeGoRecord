package leetcode

func decodeMessage(key string, message string) string {
	record := make(map[byte]byte)

	tmp := byte('a')
	for _, char := range key {
		c := byte(char)
		if c != ' ' && record[c] == 0 {
			record[c] = tmp
			tmp++
		}
	}

	res := []byte(message)

	for i := range res {
		if res[i] != ' ' {
			res[i] = record[res[i]]
		}
	}

	return string(res)
}
