// filecopy.go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	n, _ := CopyFile(
		"./eBook/examples/chapter_12/target.txt",
		"./eBook/examples/chapter_12/source.txt")
	fmt.Printf("n=%v\n", n)
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
