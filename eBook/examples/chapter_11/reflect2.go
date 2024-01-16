// reflect2.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	//reflect.ValueOf() 函数通过传递一个 x 的拷贝创建了 v，那么 v 的改变并不能更改原始的 x。
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())
	fmt.Println("type of v:", v.Type())

	//要想 v 的更改能作用到 x，那就必须传递 x 的地址
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	v = v.Elem()
	fmt.Printf("v.Elem() 的类型：%T\n", v)
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)
	fmt.Println(x)
}

/* Output:
settability of v: false
type of v: *float64
settability of v: false
The Elem of v is:  <float64 Value>
settability of v: true
3.1415
<float64 Value>
*/
