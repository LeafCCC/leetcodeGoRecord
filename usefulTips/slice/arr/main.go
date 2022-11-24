package main

import "fmt"

func main() {
	//声明一个数组
	//arr1的类型为 *[]int
	//arr2的类型为  []int
	//arr3是用数组常量的方法初始化
	var arr1 = new([2]int)
	var arr2 [2]int
	var arr3 = [3]int{1: 5, 2: 6}

	arr1[0] = 2
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("idx is %d and val is %d \n", i, arr1[i])
	}

	arr1 = &arr2
	arr2[0] = 1
	for idx, val := range arr1 {
		fmt.Printf("idx is %d and val is %d \n", idx, val)
	}

	for idx, val := range arr3 {
		fmt.Printf("idx is %d and val is %d \n", idx, val)
	}

}
