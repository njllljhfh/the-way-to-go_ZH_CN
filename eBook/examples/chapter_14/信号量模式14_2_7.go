package main

import "fmt"

func main() {
	type Empty interface{}
	var empty Empty
	N := 10

	data := make([]float64, N)
	res := make([]float64, N)
	sem := make(chan Empty, N)

	for i := 0; i < N; i++ {
		data[i] = float64(i)
	}
	fmt.Printf("data: %v\n", data)

	//for 循环的每一个迭代是并行完成的
	for i, xi := range data {
		go func(i int, xi float64) {
			res[i] = doSomething(i, xi)
			sem <- empty
		}(i, xi)
	}

	// wait for goroutines to finish
	for i := 0; i < N; i++ {
		<-sem
	}
	fmt.Printf("res: %v\n", res)
}

func doSomething(i int, xi float64) (res float64) {
	res = xi * xi
	return
}
