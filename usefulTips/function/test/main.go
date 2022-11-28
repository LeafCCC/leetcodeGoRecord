package main

import (
	"fmt"
	"time"
)

func main() {
	//输出代码的执行时间
	a := 0

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		a++
	}
	end := time.Now()
	fmt.Println("Running time is", end.Sub(start))
}
