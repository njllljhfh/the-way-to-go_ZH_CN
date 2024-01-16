package main

import "fmt"

func main() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)                     // output: [1 2 3 0 0 0 0 0 0 0]
	fmt.Printf("Copied %d elements\n", n) // n == 3
	fmt.Println("-----------------------------------")

	sl3 := []int{1, 2, 3}
	fmt.Printf("sl3 的地址：%p\n", sl3)

	sl3 = append(sl3, 4, 5, 6)      // append 方法总是返回成功，除非系统内存耗尽了。
	fmt.Printf("sl3 的地址：%p\n", sl3) // sl3 地址变了
	fmt.Println(sl3)                // output: [1 2 3 4 5 6]

	sl4 := []int{7, 8, 9}
	sl3 = append(sl3, sl4...)
	fmt.Printf("sl3 的地址：%p\n", sl3) // sl3 地址变了
	fmt.Println(sl3)
}
