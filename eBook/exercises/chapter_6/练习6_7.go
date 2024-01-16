package main

import (
	"fmt"
	"strings"
)

/*
包 strings 中的 Map 函数和 strings.IndexFunc() 一样都是非常好的使用例子。
请学习它的源代码并基于该函数书写一个程序，要求将指定文本内的所有非 ASCII 字符替换成 ? 或空格。
您需要怎么做才能删除这些字符呢？
*/

func main() {
	// 示例用法
	text := "Hello, 世界!"
	result := strings.Map(
		func(r rune) rune {
			if r > 127 {
				// 将非 ASCII 字符替换成 '?'
				//return '?'
				return -1 // 返回 负值 会删除该字符
			}
			return r
		},
		text)

	fmt.Println(result)
}
