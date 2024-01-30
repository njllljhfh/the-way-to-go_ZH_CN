package main

import "fmt"

type MyData struct {
	i int
}

func main() {
	md := MyData{i: 10}
	fmt.Printf("md.i = %v\n", md.i)

	md.SetI(5)
	fmt.Printf("md.i = %v\n", md.i)

	(&md).SetI(20)
	fmt.Printf("md.i = %v\n", md.i)

}

// 改变main中md的值
func (md *MyData) SetI(i int) {
	md.i = i
}

// 不改变main中md的值
//func (md MyData) SetI(i int) {
//	md.i = i
//}
