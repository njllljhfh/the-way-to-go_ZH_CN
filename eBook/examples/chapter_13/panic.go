package main

import "fmt"

func main() {
	fmt.Println("Starting the program")
	defer fmt.Printf("defer1\n")
	defer fmt.Printf("defer2\n")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
