// reflect1.go
// blog: Laws of Reflection
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)

	fmt.Println("type:", v.Type())
	fmt.Printf("v.Type() 的类型：%T\n", v.Type())

	fmt.Println("kind:", v.Kind())
	fmt.Printf("v.Kind() 的类型：%T\n", v.Kind())

	/*
		x 是一个 float64 类型的值，reflect.ValueOf(x).Float() 返回这个 float64 类型的实际值；
		同样的适用于 Int(), Bool(), Complex(), String()
	*/
	fmt.Println("value:", v.Float())
	fmt.Printf("v.Float() 的类型：%T\n", v.Float())

	fmt.Println(v.Interface())
	fmt.Printf("v.Interface() 的类型：%T\n", v.Interface())

	fmt.Printf("value is %5.2e\n", v.Interface())

	//y := v.Interface().(float64)
	//fmt.Println(y)
	if y, ok := v.Interface().(float64); ok {
		fmt.Printf("y = %v", y)
	} else {
		fmt.Println("v 不是 float64 类型")
	}
}

/* output:
type: float64
value: <float64 Value>
type: float64
kind: float64
value: 3.4
3.4
value is 3.40e+00
3.4
*/
