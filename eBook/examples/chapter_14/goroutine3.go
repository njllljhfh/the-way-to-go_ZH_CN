package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}

	//使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭
	//for input := range ch {
	//	fmt.Printf("%s ", input)
	//
	//}
}

// Washington Tripoli London Beijing Tokio
