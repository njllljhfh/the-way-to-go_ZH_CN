package main

import (
	"./structPack"
	"fmt"
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.
	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)
	fmt.Printf("struct1 的类型是： %T\n", struct1)
}

// Mi1 = 10
// Mf1 = 16.000000
