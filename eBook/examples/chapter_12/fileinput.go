package main

import (
	"bufio"
	"fmt"
	"io"
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

	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return // error or EOF
		}
	}
}
