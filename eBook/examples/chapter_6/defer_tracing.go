package main

import "fmt"

func trace641(s string)   { fmt.Println("entering:", s) }
func untrace641(s string) { fmt.Println("leaving:", s) }

func a641() {
	trace641("a")
	defer untrace641("a")
	fmt.Println("in a")
}

func b641() {
	trace641("b")
	defer untrace641("b")
	fmt.Println("in b")
	a641()
}
func main() {
	b641()
}
