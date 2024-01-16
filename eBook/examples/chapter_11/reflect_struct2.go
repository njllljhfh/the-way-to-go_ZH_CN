// reflect_struct2.go
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()

	typeOfT := s.Type()
	fmt.Printf("typeOfT 的类型：%T\n", typeOfT)

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	//应用反射，即使在同一个包里，也只能修改导出的字段（即大写字母开头的字段）
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}

/* Output:
0: A int = 23
1: B string = skidoo
t is now {77 Sunset Strip}
*/
