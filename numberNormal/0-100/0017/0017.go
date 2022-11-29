package leetcode

var res []string

var rcd map[byte]string = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func f(digits string, now int, tmp []byte) {
	if now == len(digits) {
		res = append(res, string(tmp))
		return
	}

	for i := range rcd[digits[now]] {
		tmp = append(tmp, rcd[digits[now]][i])
		f(digits, now+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return []string{}
	}

	f(digits, 0, []byte{})
	return res
}
