package main

import (
	"os"
)

func main() {
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile("./eBook/examples/chapter_12/test", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	f.WriteString("hello, world in a file\n")
	//fmt.Fprintf(f, "Some test data.\n")
}
