package main

import (
	"fmt"
	"myDataStructure/stack"
)

func main() {

	a := stack.Stack{}
	a.Push(1)
	a.Push(2)
	a.Push(3)
	fmt.Println(a)

	for {
		val, err := a.Pop()
		if err != nil {
			break
		}
		fmt.Println(val)
	}

}
