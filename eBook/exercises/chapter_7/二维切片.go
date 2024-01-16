package main

import "fmt"

/*
7.2.5 多维 切片
和数组一样，切片通常也是一维的，但是也可以由一维组合成高维。
通过分片的分片（或者切片的数组），长度可以任意动态变化，所以 Go 语言的多维切片可以任意切分。
而且，内层的切片必须单独分配（通过 make 函数）。
*/

func main() {
	// 创建一个二维切片
	rows, cols := 3, 4
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// 初始化二维切片的值
	value := 1
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = value
			value++
		}
	}

	// 打印二维切片
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
