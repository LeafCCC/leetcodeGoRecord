package leetcode

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	for len(a) > len(b) {
		b = "0" + b
	}

	count := 0
	res := make([]byte, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		val := int(a[i]) + int(b[i]) - 2*int('0') + count
		switch val {
		case 0:
			res[i] = '0'
			count = 0
		case 1:
			res[i] = '1'
			count = 0
		case 2:
			res[i] = '0'
			count = 1
		case 3:
			res[i] = '1'
			count = 1
		}
	}

	if count == 1 {
		res = append([]byte{'1'}, res...)
	}

	return string(res)
}
