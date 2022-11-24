package main

import "fmt"

//格式化输出的字符表示
func main() {
	a := 12
	//值的默认表示
	fmt.Printf("%v \n", a)
	//十进制表示
	fmt.Printf("%d \n", a)
	//二进制表示
	fmt.Printf("%b \n", a)
	//输出类型
	fmt.Printf("%T \n", a)
}
