package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var inputFile *os.File
	// var inputError, readerError os.Error
	// var inputReader *bufio.Reader
	// var inputString string

	inputFile, inputError := os.Open("./eBook/examples/chapter_12/input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		fmt.Printf("inputError: %v\n", inputError)
		return // exit the function on error
	}

	defer func() {
		if err := inputFile.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	inputReader := bufio.NewReader(inputFile)

	//2) 带缓冲的读取
	buf := make([]byte, 1024)
	n, readerError := inputReader.Read(buf)
	if readerError != nil {
		fmt.Printf("读取错误：%v\n", readerError)
	}
	fmt.Printf("读取到的字节数: %d\n", n)
	fmt.Printf("文件内容: \n%s", buf)
}
