// panic_recover.go
package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
	//fmt.Println("good end")
}

func test() (err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
	return
}

func main() {
	fmt.Printf("Calling test\r\n")
	err := test()
	if err != nil {
		fmt.Printf("tes函数调用失败：%v\n", err)
	}
	fmt.Printf("Test completed\r\n")
}

/* Output:
Calling test
Panicing bad end
Test completed
*/
