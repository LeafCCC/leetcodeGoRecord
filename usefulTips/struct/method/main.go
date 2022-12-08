package main

import "fmt"

type Account struct {
	income  int
	outcome int
}

// 改变输入输出并返回目前净收入
// 要改变struct内的值必须使用指针 *Account
func (this *Account) change(in int, out int) int {
	this.income += in
	this.outcome += out
	return this.income - this.outcome
}

//指针方法和值方法都可以在指针或非指针上被调用

func main() {
	//结构体指针
	a := new(Account)
	resa := a.change(1, 2)
	fmt.Println(a.income, a.outcome, resa)

	//结构体
	b := Account{}
	resb := b.change(3, 1)
	fmt.Println(b.income, b.outcome, resb)
}
