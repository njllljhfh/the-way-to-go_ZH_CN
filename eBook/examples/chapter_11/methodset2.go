// methodset2.go
package main

import (
	"fmt"
)

/*
在第 10.6.3 节及例子 methodset1.go 中我们看到，作用于变量上的方法实际上是不区分变量到底是指针还是值的。
当碰到接口类型值时，这会变得有点复杂，原因是接口变量中存储的具体值是不可寻址的，
幸运的是，如果使用不当编译器会给出错误。
*/

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in function argument:
	// List does not implement Appender (Append method requires pointer receiver)
	// CountInto(lst, 1, 10) // INVALID: Append has a pointer receiver

	if LongEnough(lst) { // VALID: Identical receiver type
		fmt.Printf(" - lst is long enough")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) // VALID: Identical receiver type
	if LongEnough(plst) {  // VALID: a *List can be dereferenced for the receiver
		fmt.Printf(" - plst is long enough") //  - plst is long enoug
	}
}
