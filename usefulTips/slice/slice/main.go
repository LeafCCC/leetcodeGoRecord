// 切片在内存中的组织方式实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度len以及切片容量cap
package main

import (
	"fmt"
	"sort"
)

func doSomething(s []int) {
	s[0] = -1
}

func main() {
	//初始化一个切片 并直接赋值
	var slice1 = []int{1, 2}
	//通过一个数组得到切片
	var arr [5]int
	var slice2 = arr[1:3]
	//通过make初始化切片
	//make和new的区别
	//new(T)为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的
	//new确定分配内存的大小 如数组和结构体
	//make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型
	//make创造只有在运行过程中才能确定内存大小的数据 如map 、slice 、channel
	//func make([]T, len, cap)
	slice4 := make([]int, 5)

	for idx, val := range slice1 {
		fmt.Printf("idx is %d and val is %d \n", idx, val)
	}

	fmt.Printf("slice1 len is %d and cap is %d \n", len(slice1), cap(slice1))

	for idx, val := range slice2 {
		fmt.Printf("idx is %d and val is %d \n", idx, val)
	}

	fmt.Printf("slice2 len is %d and cap is %d \n", len(slice2), cap(slice2))

	slice3 := slice2
	slice3[1] = 100
	for idx, val := range slice2 {
		fmt.Printf("idx is %d and val is %d \n", idx, val)
	}

	doSomething(slice2)

	for idx, val := range slice2 {
		fmt.Printf("idx is %d and val is %d \n", idx, val)
	}

	fmt.Printf("slice4 len is %d and cap is %d \n", len(slice4), cap(slice4))

	slice4 = slice4[2:4]

	//截断前面的部分会导致cap降低
	fmt.Printf("slice4 len is %d and cap is %d \n", len(slice4), cap(slice4))

	//slice的复制
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n)

	//slice的追加
	sl := []int{1, 2, 3}
	sl = append(sl, 4)
	fmt.Println(sl)
	sl = append(sl, sl_from...)
	fmt.Println(sl)

	//排序
	test1 := []int{5, 4, 8, 7, 9, 1, 2, 3, 1}
	sort.Ints(test1)
	fmt.Println(test1)
	sort.Sort(sort.Reverse(sort.IntSlice(test1))) //倒序
	fmt.Println(test1)

	test2 := []string{"abc", "a", "b", "A", "B", "ABC", "ab"}
	sort.Strings(test2)
	fmt.Println(test2)

	//二维切片
	const q = 2
	dp1 := make([][q]int, q)
	fmt.Println(dp1)

	//如果不用常量来定义
	n = 2
	dp2 := make([][]int, n)
	for i := range dp2 {
		dp2[i] = make([]int, n)
	}
	fmt.Println(dp2)

	//二维切片的排序
	var intervals = [][]int{{2, 3}, {1, 2}, {4, 5}, {-1, 4}}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)

	//删除某元素
	deleteT := []int{1, 2, 3, 4, 5}
	deleteT = append(deleteT[:4], deleteT[5:]...)
	fmt.Println(deleteT)
}
