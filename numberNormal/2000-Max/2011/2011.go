package leetcode

func finalValueAfterOperations(operations []string) int {
	rcd := map[string]int{
		"++X": 1,
		"X++": 1,
		"--X": -1,
		"X--": -1,
	}

	x := 0

	for _, i := range operations {
		x += rcd[i]
	}

	return x
}
