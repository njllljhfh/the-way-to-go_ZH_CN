package main

/*
练习 6.9 不使用递归但使用闭包改写第 6.6 节中的斐波那契数列程序。
*/

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	// println(f(), f(), f(), f(), f())
	for i := 0; i <= 9; i++ {
		println(i+2, f())
	}
}
