package main

import (
	"bytes"
	"fmt"
)

// 使用Buffer拼接字符串更快更节省内存
func main() {
	var buffer bytes.Buffer

	a := []byte{65, 66, 125, 126}
	buffer.Write(a)

	fmt.Print(buffer.String(), "\n")
}
