package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * 3.1415926
}

func main() {

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	c := Circle{1.2}
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q, c}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	//类型断言 用来检测一个接口具体包含的类型是什么
	if i, ok := shapes[0].(Rectangle); ok {
		fmt.Printf("%T", i)
		fmt.Println(i, ok)
	}

	//类型判断 switch-type
	//switch shapes[1].(type) 不用到具体值时可以这么写
	switch t := shapes[1].(type) {
	case *Square:
		fmt.Printf("The type is %T and the val is %v", t, t)
	case Rectangle:
		fmt.Printf("The type is %T and the val is %v", t, t)
	case Circle:
		fmt.Printf("The type is %T and the val is %v", t, t)
	case nil:
		fmt.Printf("None")
	default:
		fmt.Printf("Unexpected type %T and val %v", t, t)
	}

	// 	Go 语言规范定义了接口方法集的调用规则：
	//
	// 			类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
	// 			类型 T 的可调用方法集包含接受者为 T 的所有方法
	//			类型 T 的可调用方法集不包含接受者为 *T 的方法(此处注意)

}
