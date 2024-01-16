package main

import "fmt"

/*
练习 6.5

使用递归函数从 10 打印到 1。
*/

func main() {
	printrec(1)
}

func printrec(i int) {
	if i > 10 {
		return
	}
	printrec(i + 1)
	fmt.Printf("%d ", i)
}
