// Q20_linked_list.go
package main

import (
	"container/list"
	"fmt"
)

/*
练习 9.1

使用 container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。
*/

func main() {
	lst := list.New()
	lst.PushBack(100)
	lst.PushBack(101)
	lst.PushBack(102)
	//fmt.Println("Here is the double linked list:\n", lst)
	for e := lst.Front(); e != nil; e = e.Next() {
		// fmt.Println(e)
		fmt.Println(e.Value)
	}
}

/* Example output:
&{0x12542bc0 <nil> 0x12547590 1}
&{0x12542ba0 0x12542be0 0x12547590 2}
&{<nil> 0x12542bc0 0x12547590 4}

100
101
102
*/
