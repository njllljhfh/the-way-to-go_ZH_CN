package main

import (
	"fmt"
	"time"
)

var num1 int
var num2 int

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
			num1 += 1
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
			num2 += 1
		}
	}
}
