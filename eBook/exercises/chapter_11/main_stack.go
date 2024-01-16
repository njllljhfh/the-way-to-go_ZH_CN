// main_stack.go
package main

import (
	"./stack"
	"fmt"
)

var st1 stack.Stack

func main() {
	fmt.Printf("main 内 st1的地址：%p\n", &st1)

	st1.Push("Brown")
	_, err := st1.Top() // Top内 &stack 的地址，与 main函数体内 &st1 的地址不一样
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	st1.Push(3.14)
	st1.Push(100)
	st1.Push([]string{"Java", "C++", "Python", "C#", "Ruby"})
	for {
		item, err := st1.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}

/* Output:
[Java C++ Python C# Ruby]
100
3.14
Brown
*/
