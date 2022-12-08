package main

import "fmt"

type Base struct {
	id int
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(newId int) {
	b.id = newId
}

type Person struct {
	Base
	FirstName, LastName string
}

type Employee struct {
	Person
	salary int
}

func main() {
	e1 := &Employee{Person{Base{1}, "Leaf", "CCC"}, 12}
	fmt.Println(e1.Id())
	e1.SetId(3)
	fmt.Println(e1.Id())
}
