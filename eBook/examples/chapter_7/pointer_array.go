package main

import "fmt"

func f(a [3]int) { fmt.Println(a) }
func fp(a *[3]int) {
	fmt.Println(a)
	fmt.Printf("a 的类型： %T\n", a)
}

func main() {
	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar
}
