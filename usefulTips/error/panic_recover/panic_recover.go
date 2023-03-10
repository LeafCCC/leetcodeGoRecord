package main

import (
	"fmt"
	"time"
)

func main() {
	// 第一个协程
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("Error: ", err)
			}
		}()
		var i int
		for {
			i++
			fmt.Println("协程1")
			time.Sleep(1 * time.Second)
			// 3秒后发生panic
			if i == 3 {
				panic("协程1异常退出")
			}
		}

	}()
	// 第二个协程
	go func() {
		for {
			fmt.Println("协程2")
			time.Sleep(1 * time.Second)
		}
	}()
	// 让主协程不退出
	for {
		time.Sleep(1 * time.Second)
	}
}
