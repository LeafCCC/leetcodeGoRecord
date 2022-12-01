package leetcode

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	rcd := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []byte{}

	t := []byte(s)

	for _, i := range t {
		if _, ok := rcd[i]; ok {
			stack = append(stack, i)
		} else {
			if len(stack) != 0 && rcd[stack[len(stack)-1]] == i {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
