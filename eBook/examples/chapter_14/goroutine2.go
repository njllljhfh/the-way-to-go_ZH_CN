package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	fmt.Println("11111")
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	fmt.Println("开始发送")
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	fmt.Println("\n发送结束")
}

func getData(ch chan string) {
	fmt.Println("开始接收")
	var input string
	// time.Sleep(1e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
	fmt.Println("停止接受")
}

// Washington Tripoli London Beijing Tokio
