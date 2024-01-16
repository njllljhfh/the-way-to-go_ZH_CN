package main

import "fmt"

func x(a []int) {
	for i, _ := range a {
		a[i] = i + 1
	}
}

func y(a []int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	//下面两种方法可以生成相同的切片:
	//a := new([10]int)[0:5]
	a := make([]int, 5, 10) //推荐
	fmt.Printf("a 的类型：%T\n", a)
	fmt.Println("-----------------------------------------------------")

	y(a)
	fmt.Println("-----------------------------------------------------")
	x(a)
	y(a)
	fmt.Println("-----------------------------------------------------")
}
