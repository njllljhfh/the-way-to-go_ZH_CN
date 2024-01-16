package main

import "fmt"

/*
练习 6.4

重写本节中生成斐波那契数列的程序并返回两个命名返回值（详见第 6.2 节），
即数列中的位置和对应的值，例如 5 与 4，89 与 10。
*/

func main() {
	pos := 4
	result, pos := fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, result)
	pos = 10
	result, pos = fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, result)
}

func fibonacci(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		v1, _ := fibonacci(n - 1)
		v2, _ := fibonacci(n - 2)
		val = v1 + v2
	}
	pos = n
	return
}
