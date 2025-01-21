package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	// os.Args is slice of strings the contains command line args passed while running a Go program
	// os.Args[0] is the name of the program itself
	// os.Args[1] is the first arg passed after the program's name
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Something went wrong while reading", os.Args[1])
		os.Exit(1)
	}

	// This works because "File" also implements the *Reader* interface
	io.Copy(os.Stdout, f)

	// #1. Output contents of a file
	// go run main.go myFile.txt
	// #2. Output contents of source code file itself
	// go build main.go
	// ./main main.go
}
