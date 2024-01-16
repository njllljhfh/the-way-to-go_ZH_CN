package main

import (
	"fmt"
	"os"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func main() {
	s := FindDigits("D:\\code\\github\\the-way-to-go_ZH_CN\\eBook\\exercises\\chapter_7\\切片和垃圾回收7_6_8.txt")
	fmt.Println(s)

	s2 := FindFileDigits("D:\\code\\github\\the-way-to-go_ZH_CN\\eBook\\exercises\\chapter_7\\切片和垃圾回收7_6_8.txt")
	fmt.Println(s2)

}

func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// 事实上，上面这段代码只能找到第一个匹配正则表达式的数字串。要想找到所有的数字，可以尝试下面这段代码：
func FindFileDigits(filename string) []byte {
	fileBytes, _ := os.ReadFile(filename)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	c := make([]byte, 0)
	//fmt.Printf("c 的内存地址：%p\n", c)
	for _, bytes := range b {
		c = append(c, bytes...)
		//fmt.Printf("c 的内存地址：%p\n", c)
	}
	return c
}
