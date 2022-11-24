package leetcode

func lengthOfLongestSubstring(s string) int {
	//key记录字符 val记录其最后一次出现的位置
	rcd := make(map[byte]int)

	begin := 0
	res := 0
	for now, _ := range s {
		if idx, ok := rcd[s[now]]; ok {
			if idx >= begin {
				if now-begin > res {
					res = now - begin
				}
				begin = idx + 1
			}
		}
		rcd[s[now]] = now
	}
	if len(s)-begin > res {
		res = len(s) - begin
	}
	return res
}
