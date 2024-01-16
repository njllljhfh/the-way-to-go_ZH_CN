package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	//6.10 使用闭包调试
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	fmt.Println(myFactorial(uint64(4)))
	where()

	fmt.Println(myFactorial(uint64(5)))
	where()
}

func myFactorial(num uint64) (res uint64) {
	if num <= 0 {
		res = 1
	} else {
		res = num * myFactorial(num-1)
	}
	return
}
