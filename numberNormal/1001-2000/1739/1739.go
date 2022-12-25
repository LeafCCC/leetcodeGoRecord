package leetcode

func minimumBoxes(n int) int {
	//每一层的 a[n] = n² + n / 2
	//总的 Sn = 2n³ + 6n² + 4n / 12

	S := func(n int) int {
		return (2*n*n*n + 6*n*n + 4*n) / 12
	}
	a := func(n int) int {
		return (n*n + n) / 2
	}

	t := 1
	for S(t) <= n {
		t++
	}
	t--

	remain := n - S(t)

	if remain == 0 {
		return a(t)
	}

	m := 1
	for a(m) < remain {
		m++
	}

	return a(t) + m
}
