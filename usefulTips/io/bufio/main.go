package main

import (
	"bufio"
	"fmt"
	"os"
)

// var inputReader *bufio.Reader

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("Hello, %s!", input)
	}
}
