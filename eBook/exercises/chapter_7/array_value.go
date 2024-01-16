package main

import "fmt"

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	arr2 := arr1 // 当数组赋值时，发生了数组内存拷贝。
	arr2[2] = 100

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array arr1 at index %d is %d\n", i, arr1[i])
	}
	fmt.Println()
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Array arr2 at index %d is %d\n", i, arr2[i])
	}
}
