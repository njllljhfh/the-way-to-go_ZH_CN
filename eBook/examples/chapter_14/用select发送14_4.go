package main

import (
	"fmt"
	"time"
)

func sendData(ch1, ch2 chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		select {
		case ch1 <- i:
			fmt.Println("Sent to ch1:", i)
		case ch2 <- i * 2:
			fmt.Println("Sent to ch2:", i*2)
		}
	}
	close(ch1)
	close(ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendData(ch1, ch2)

	ch1IsClosed := false
	ch2IsClosed := false
	for {
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println("Received from ch1:", data)
			} else {
				fmt.Println("ch1 is closed")
				ch1IsClosed = true
			}
		case data, ok := <-ch2:
			if ok {
				fmt.Println("Received from ch2:", data)
			} else {
				fmt.Println("ch2 is closed")
				ch2IsClosed = true
			}
		}

		if ch1IsClosed && ch2IsClosed {
			break
		}
	}
}
