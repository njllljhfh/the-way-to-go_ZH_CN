package main

import (
	"./trans"
	"fmt"
)

var twoPi1 = 2 * trans.Pi

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi1) // 2*Pi = 6.283185307179586
	trans.X1()                        //测试init函数
}
