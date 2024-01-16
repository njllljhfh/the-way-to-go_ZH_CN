// factorial.go
package main

import (
	"fmt"
)

/*
练习 6.6

实现一个输出前 30 个整数的阶乘的程序。

n! 的阶乘定义为：n! = n * (n-1)!, 0! = 1，因此它非常适合使用递归函数来实现。

然后，使用命名返回值来实现这个程序的第二个版本。

特别注意的是，使用 int 类型最多只能计算到 12 的阶乘，因为一般情况下 int 类型的大小为 32 位，继续计算会导致溢出错误。那么，如何才能解决这个问题呢？

最好的解决方案就是使用 big 包（详见第 9.4 节）。
*/

func main() {
	for i := uint64(0); i < uint64(30); i++ {
		fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
	}
}

/* unnamed return variables:
func Factorial(n uint64) uint64 {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}
*/

// named return variables:
func Factorial(n uint64) (fac uint64) {
	fac = 1
	if n > 0 {
		fac = n * Factorial(n-1)
		return
	}
	return
}

// int: correct results till 12!
// uint64: correct results till 21!
