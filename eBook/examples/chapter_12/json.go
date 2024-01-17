// json.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	Remark    string
}

func main() {
	filePath := "eBook/examples/chapter_12/"
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	//fmt.Printf("JSON format: %T\n", js)
	fmt.Printf("JSON format: %s\n", js)
	fmt.Println("------------------------------------------")

	//json反序列化
	/*
		虽然反射能够让 JSON 字段去尝试匹配目标结构字段；但是只有真正匹配上的字段才会填充数据。
		字段没有匹配不会报错，而是直接忽略掉。
	*/
	var v VCard
	json.Unmarshal(js, &v)
	fmt.Printf("v = %v\n", v)
	fmt.Println("------------------------------------------")

	// using an encoder:
	file, _ := os.OpenFile(filePath+"vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
