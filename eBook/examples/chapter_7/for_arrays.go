package main

import "fmt"

func main() {
	var arr1 [5]int
	//var arr1 = new([5]int)
	fmt.Printf("arr1的类型：%T\n", arr1)
	fmt.Printf("%v\n", arr1)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
}
