// leaving out the length 4 in [4] uses  slice instead of an array
// which is generally better for performance
package main

import "fmt"

/*
练习 7.7

a) 写一个 Sum 函数，传入参数为一个 32 位 float 数组成的数组 arrF，返回该数组的所有数字和。

如果把数组修改为切片的话代码要做怎样的修改？如果用切片形式方法实现不同长度数组的的和呢？

b) 写一个 SumAndAverage 方法，返回两个 int 和 float32 类型的未命名变量的和与平均值。
*/

func main() {
	var arr = [4]float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("The sum of the array is: %f\n", SumArr(arr))

	var slice = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("The sum of the slice is: %f\n", SumSlice(slice))

	var b = []int{1, 2, 3, 4, 5}
	sum, average := SumAndAverage(b)
	fmt.Printf("The sum of the slice is: %d, and the average is: %f", sum, average)
}

func SumArr(a [4]float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}

func SumSlice(a []float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}

func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum, float32(sum / len(a))
}
