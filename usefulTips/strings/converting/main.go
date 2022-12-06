package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, err1 := strconv.Atoi("-123")
	fmt.Println(a, err1)

	b := strconv.Itoa(12322222222222)
	fmt.Println(b)
}
