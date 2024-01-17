package main

import (
	"flag" // command line option parser
	"fmt"
	"os"
)

/*
切换到chapter_12，执行以下命令，查看区别
go run .\echo.go A B C
go run .\echo.go -n A B C
*/

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	// 命令行加 -n 参数时，该值为true，不加则为false
	fmt.Printf("*NewLine = %v\n", *NewLine)
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
