package main

import "fmt"

/*
练习 7.6
把一个缓存 buf 分片成两个 切片：第一个是前 n 个 bytes，后一个是剩余的，用一行代码实现。
*/

func main() {
	buf := []byte{'a', 'b', 'c', 'd', 'e', 'f'}

	n := 2
	firstSlice, remainingSlice := buf[:n], buf[n:]

	fmt.Printf("firstSlice = %v\n", firstSlice)
	fmt.Printf("remainingSlice = %v\n", remainingSlice)
}
