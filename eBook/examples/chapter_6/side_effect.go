// side_effect.go
package main

import (
	"fmt"
)

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n
	fmt.Println("Multiply:", *reply) // Multiply: 0
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Println("Multiply:", n)      // Multiply: 50
}
