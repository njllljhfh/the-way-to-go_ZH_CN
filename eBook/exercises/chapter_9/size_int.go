// size_int.go
package main

import (
	"fmt"
	"unsafe"
)

/*
练习 9.2

通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。
*/

func main() {
	var i int = 10
	size := unsafe.Sizeof(i)
	fmt.Println("The size of an int is: ", size)
}

// The size of an int is:  4
