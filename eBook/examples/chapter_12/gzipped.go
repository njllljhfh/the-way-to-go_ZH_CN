// gzipped.go
package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fName := "./eBook/examples/chapter_12/MyFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName) // Open内部以只读模式 调用了OpenFile
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	defer fi.Close()
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		fmt.Println("else")
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
