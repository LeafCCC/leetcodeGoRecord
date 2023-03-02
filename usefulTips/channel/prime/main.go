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

// 这里可以看成是一个个通道的串联 即 ch1->ch2->ch3 有点链表的意思
// 每个通道负责筛一下 能通过所有通道的输出并形成一个新的筛除通道
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
