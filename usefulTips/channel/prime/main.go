package main

import (
	"fmt"
)

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		tmp := <-in

		if tmp%prime != 0 {
			out <- tmp
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		if prime > 1000 {
			break
		}
		fmt.Printf("%d ", prime)
		nextCh := make(chan int)
		go filter(ch, nextCh, prime)
		ch = nextCh
	}
}
