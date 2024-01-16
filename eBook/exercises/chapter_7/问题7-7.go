package main

import "fmt"

/*
问题 7.7
如果 a 是一个切片，那么 a[n:n] 的长度是多少？
a[n:n+1] 的长度又是多少？
*/

func main() {
	sl := []int{10, 20, 30, 40, 50}
	fmt.Printf("sl[3:3] = %v, 长度=%d, 容量=%d\n", sl[3:3], len(sl[3:3]), cap(sl[3:3]))
	fmt.Printf("sl[3:4] = %v, 长度=%d, 容量=%d\n", sl[3:4], len(sl[3:4]), cap(sl[3:4]))
}
