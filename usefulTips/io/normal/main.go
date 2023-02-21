package main

import "fmt"

func main() {
	var name string

	fmt.Println("Please input your name:")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!", name)

}
