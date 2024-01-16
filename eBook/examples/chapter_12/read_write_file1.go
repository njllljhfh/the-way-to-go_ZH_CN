// read_write_file.go
package main

import (
	"fmt"
	"os"
)

func main() {
	inputFile := "./eBook/examples/chapter_12/products.txt"
	outputFile := "./eBook/examples/chapter_12/products_copy.txt"
	buf, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = os.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
