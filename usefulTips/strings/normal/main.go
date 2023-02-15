package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
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

	//另一种拼接 这种效率更高
	t = []byte(a)
	t = append(t, b...)
	a = string(t)
	fmt.Println(a)

	s := "  abcd   "
	//删除开头和结尾的空格
	s = strings.TrimSpace(s)
	fmt.Println(s)

	t1 := "a"
	t2 := "b"
	start := time.Now()
	for i := 0; i <= 10_0000; i++ {
		t1 = t1 + t2
	}
	fmt.Println(time.Now().Sub(start))

	t1 = "a"
	start = time.Now()
	for i := 0; i <= 10_0000; i++ {
		t1 += t2
	}
	fmt.Println(time.Now().Sub(start))

	t1 = "a"
	start = time.Now()
	t = []byte(t1)
	for i := 0; i <= 100_0000; i++ {

		t = append(t, t2...)

	}
	t1 = string(t)
	fmt.Println(time.Now().Sub(start))

	var hun = "混元形意太极拳a"
	fmt.Printf("hun 占用 %d 个字符数\n", utf8.RuneCountInString(hun))
	fmt.Printf("hun 占用 %d 个字节数\n", len([]byte(hun)))

}
