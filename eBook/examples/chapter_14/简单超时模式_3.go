package main

import "fmt"

/*
第三种形式：假设程序从多个复制的数据库同时读取。只需要一个答案，需要接收首先到达的答案，
Query 函数获取数据库的连接切片并请求。并行请求每一个数据库并返回收到的第一个响应：
*/

type Conn struct {
	index int
}

type Result struct {
	context  string
	cnnIndex int
}

func main() {
	for i := 0; i < 20; i++ {
		conns := make([]Conn, 3)
		for j := 0; j < len(conns); j++ {
			conns[j].index = j
		}
		res := Query(conns, "请求")
		fmt.Printf("res = %v\n", res)
	}
}

func Query(conns []Conn, query string) Result {
	ch := make(chan Result, 1)
	for _, conn := range conns {
		go func(c *Conn) {
			select {
			case ch <- c.DoQuery(query):
			default:
			}
		}(&conn)
	}
	return <-ch
}

func (conn *Conn) DoQuery(q string) (res Result) {
	res.context = q
	res.cnnIndex = conn.index
	return
}

func (res Result) String() string {
	return fmt.Sprintf("%s-cnnIndex: %d", res.context, res.cnnIndex)
}
