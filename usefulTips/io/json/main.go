package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []Address
}

func main() {
	ha := Address{"home", "Suzhou", "China"}
	sa := Address{"school", "Shanghai", "China"}
	myCard := VCard{"CCC", "Leaf", []Address{ha, sa}}

	//序列化
	js, _ := json.Marshal(myCard)
	fmt.Printf("js format is: %s, its type is %T \n", js, js)

	//写文件
	// file, _ := os.OpenFile("./leetcodeGoRecord/usefulTips/io/json/test.json", os.O_CREATE|os.O_WRONLY, 0666)
	file, _ := os.Create("./leetcodeGoRecord/usefulTips/io/json/test.json")

	encoder := json.NewEncoder(file)
	err := encoder.Encode(js)

	if err != nil {
		fmt.Println("sth wrong in json encoding...")
	}

	file.Close()

	//读文件
	var rec []byte
	file2, _ := os.Open("./leetcodeGoRecord/usefulTips/io/json/test.json")
	defer file2.Close()
	decoder := json.NewDecoder(file2)
	err2 := decoder.Decode(&rec)

	if err2 != nil {
		fmt.Println("sth wrong in json decoding...")
	}

	//反序列化
	var res VCard
	_ = json.Unmarshal(rec, &res)
	fmt.Println(res)

}
