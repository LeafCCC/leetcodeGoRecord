package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	Addresses []*Address
}

func main() {
	ha := &Address{"home", "Suzhou", "China"}
	sa := &Address{"school", "Shanghai", "China"}
	myCard := VCard{"CCC", "Leaf", []*Address{ha, sa}}

	js, _ := json.Marshal(myCard)
	fmt.Printf("js format is: %s, its type is %T \n", js, js)

	file, _ := os.OpenFile("./leetcodeGoRecord/usefulTips/io/json/test.json", os.O_CREATE|os.O_WRONLY, 0666)

	encoder := json.NewEncoder(file)
	err := encoder.Encode(js)

	if err != nil {
		fmt.Println("sth wrong in json encoding...")
	}

	file.Close()

	var rec []byte
	file2, _ := os.Open("./leetcodeGoRecord/usefulTips/io/json/test.json")
	r := io.Reader(file2)
	defer file2.Close()
	decoder := json.NewDecoder(r)
	err2 := decoder.Decode(&rec)

	if err2 != nil {
		fmt.Println("sth wrong in json decoding...")
	}

	var res VCard
	_ = json.Unmarshal(rec, &res)
	fmt.Println(res)

}
