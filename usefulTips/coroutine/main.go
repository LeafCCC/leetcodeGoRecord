package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	//go getData(ch)

	fmt.Println(<-ch)

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
