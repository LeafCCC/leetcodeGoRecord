package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//如果不知道行数 和每行的元素数量 如何读取？
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() { //每次读一行
		data := strings.Split(input.Text(), " ")
		fmt.Println(data)
	}
}
