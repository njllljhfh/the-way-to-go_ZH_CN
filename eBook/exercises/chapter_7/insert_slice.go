// insert_slice.go
package main

import (
	"fmt"
)

/*
练习 7.11

写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。
*/

func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := InsertStringSlice(s, in, 0) // at the front
	fmt.Println(res)                   // [A B C M N O P Q R]
	res = InsertStringSlice(s, in, 3)  // [M N O A B C P Q R]
	fmt.Println(res)
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}
