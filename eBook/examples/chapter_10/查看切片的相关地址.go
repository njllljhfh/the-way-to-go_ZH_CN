package main

import "fmt"

func main() {
	a := make([]int, 10)
	fmt.Println(a)
	fmt.Printf("&a的地址：%p\n", &a)
	fmt.Printf("a[0]的地址：%p\n", &(a[0]))
	fmt.Printf("a的地址：%p\n", a)
	fmt.Println("------------------------------------")

	a = a[2:]
	fmt.Printf("&a的地址：%p\n", &a)
	fmt.Printf("a[0]的地址：%p\n", &(a[0]))
	fmt.Printf("a的地址：%p\n", a)
}
