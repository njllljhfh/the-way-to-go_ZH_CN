package main

import "fmt"

/*
问题 7.6 通过使用省略号操作符 ... 来实现累加方法。
*/

func main() {
	sum1 := sum76(10, 20, 30, 40, 50)
	fmt.Println(sum1)
}

func sum76(nums ...int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}
