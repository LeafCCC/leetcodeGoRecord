package main

import "fmt"

func main() {
	//两种初始化方式
	record := make(map[string]int)
	record["a"]++ //这样也可以 即使key不存在 则为默认值

	record2 := map[string]int{"b": 2}
	fmt.Println(record["a"], record2["b"])

	//事先申明好大小可以提高性能
	map3 := make(map[string]float32, 100)
	map3["c"] = 3
	map3["d"] = 4

	//判断键值是否存在
	if val, isPresent := map3["c"]; isPresent {
		fmt.Println(val)
		delete(map3, "c") //删除键值
	}

	map3["e"] = 5
	//遍历map
	for key, val := range map3 {
		fmt.Println("key is", key, ", val is", val)
	}

	fmt.Println(len(map3))

}
