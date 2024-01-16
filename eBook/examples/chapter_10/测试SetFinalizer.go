package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type myInt int

func main() {
	doSomething()
	runtime.GC()
	time.Sleep(500 * time.Millisecond)

	//查看内存
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}

func (i myInt) String() (str string) {
	str = strconv.Itoa(int(i))
	return
}

func doSomething() {
	i := myInt(10)
	fmt.Println(i)
	runtime.SetFinalizer(&i, func(obj *myInt) { fmt.Printf("%v 的内存被回收了\n", *obj) })
}
