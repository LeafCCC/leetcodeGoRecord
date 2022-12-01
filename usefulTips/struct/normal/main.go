package main

import "fmt"

//定义一个结构体
type student struct {
	id   int
	name string
}

func main() {
	//初始化一个结构体
	a := &student{id: 1, name: "Alice"} //类型为*student
	b := student{2, "Bob"}              //类型为student
	c := new(student)                   //类型为*student

	//无论是指针类型 还是普通类型 赋值过程类似
	c.id = 3
	c.name = "Cat"
	fmt.Println(a, b, c)
	b.name = "NewBob"
	fmt.Println(a, b, c)
}
