package main

import (
	"fmt"
	"strconv"
)

type myInt int

func main() {
	var num myInt
	num = 6
	fmt.Printf("num = %s\n", &num)

	var a any
	//a = num  //num没有实现String
	a = &num //&num实现了String

	if v, ok := a.(fmt.Stringer); ok {
		fmt.Printf("v = %v\n", v)
	}
}

func (i *myInt) String() string {
	return strconv.Itoa(int(*i))
}
