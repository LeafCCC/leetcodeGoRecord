package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	ch <- "1"
	fmt.Println(<-ch)

	func(x string) {
		fmt.Println(x)
	}("test")

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "1"
	ch <- "2"
	ch <- "3"
}

func getData(ch chan string) {
	time.Sleep(0.99e9)
	for {
		s := <-ch
		fmt.Println(s)
	}
}
