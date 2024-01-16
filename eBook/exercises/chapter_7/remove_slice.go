// remove_slice.go
package main

import (
	"fmt"
)

/*
练习 7.12

写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除。
*/

func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	res := RemoveStringSlice(s, 2, 4)
	fmt.Println(res) // [M N Q R]
}

func RemoveStringSlice(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}
