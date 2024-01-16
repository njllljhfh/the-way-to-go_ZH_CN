package main

import "fmt"

/*
练习 7.5
给定切片 sl，将一个 []byte 数组追加到 sl 后面。
写一个函数 Append(slice, data []byte) []byte，该函数在 sl 不能存储更多数据的时候自动扩容。
*/

func Append(slice, data []byte) []byte {
	if len(slice)+len(data) > cap(slice) {
		newCap := (len(slice) + len(data)) * 2
		newSlice := make([]byte, len(slice), newCap)
		//fmt.Printf("slice 的类型：%T\n", slice)
		//fmt.Printf("slice 的地址：%p\n", slice)
		//fmt.Printf("newSlice 的类型：%T\n", newSlice)
		//fmt.Printf("newSlice 的地址：%p\n", newSlice)
		copy(newSlice, slice)
		slice = newSlice
		//fmt.Printf("slice 的类型：%T\n", slice)
		//fmt.Printf("slice 的地址：%p\n", slice)
	}
	slice = append(slice, data...)
	return slice
}

func main() {
	originalSlice := []byte{1, 2, 3}
	dataToAppend := []byte{4, 5, 6}

	result := Append(originalSlice, dataToAppend)
	fmt.Println(result)
}
