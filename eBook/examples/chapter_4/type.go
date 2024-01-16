package main

import "fmt"

// 定义 TZ 类型
type TZ int

// 给 int 定义一个别名
type intAlias = int

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value: %d\n", c)
	fmt.Printf("c 的类型：%T\n", c)

	var d intAlias
	fmt.Printf("d 的类型：%T\n", d)

	//fmt.Println(c == d)  //报错 invalid operation: c == d (mismatched types TZ and int)
}
