package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abcdefg"
	b := "hijklmn"

	//修改某个字符
	t := []byte(a)
	t[0] = 'z'
	a = string(t)
	fmt.Println(a)

	//拼接字符串
	a = a + b
	fmt.Println(a)

	s := "  abcd   "
	//删除开头和结尾的空格
	s = strings.TrimSpace(s)
	fmt.Println(s)
}
