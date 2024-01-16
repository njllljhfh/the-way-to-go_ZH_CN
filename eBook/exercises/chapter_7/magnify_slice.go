package main

import "fmt"

/*
练习 7.9

给定一个slices []int 和一个 int 类型的因子factor，扩展 s 使其长度为 len(s) * factor。
*/

var s []int

func main() {
	s = []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(s))
	fmt.Println(s)
	s = enlarge(s, 5)
	fmt.Println("The length of s after enlarging is:", len(s))
	fmt.Println(s)
}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	// fmt.Println("The length of ns  is:", len(ns))
	copy(ns, s)
	//fmt.Println(ns)
	s = ns
	//fmt.Println(s)
	//fmt.Println("The length of s after enlarging is:", len(s))
	return s
}
