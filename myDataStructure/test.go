package main

import (
	"fmt"
	"myDataStructure/myStack"
)

func main() {
	//stack
	a := myStack.Stack{}
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

	// h := &intHeap{5, 4, 2, 3, 1}
	// fmt.Println(*h)
	// heap.Init(h)
	// fmt.Println(*h)
}
