package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// ** Reading from a file
	// f, err := os.Open("read.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// Creates a stream of bytes; as if they were coming from a network or file
	f := strings.NewReader("Hello")

	n, err := readBytes(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Letters read from the file: %d\n", n)

	// ** Write to a file
	f1, err := os.Create("write.txt")
	if err != nil {
		panic(err)
	}

	defer f1.Close()

	n1, _ := writeBytes("Hello world", f1)
	fmt.Printf("Bytes written: %d\n", n1)
}

// io.Reader
// - An interface that has one method `Read`
// - Read(p []byte) (n int, err error)
// - It returns the no. of bytes processed and an error
// - Takes a `byte` slice to read the data into
// - Common implementation
//   - `*os.File`: used to read data from a file
//   - `strings.Reader`: creates a new reader from the provided string
//   - `bufio.Reader`: implements Read for improved performance, using buffers for reading
//   - `http.Request.Body`: allows you to supply the request body via Reader
//   - `bytes.Buffer`: used for reading from in-memory bytest buffer
func readBytes(r io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)

	for {
		n, err := r.Read(buffer)
		for _, ch := range buffer[:n] {
			if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
				count++
			}
		}

		// var EOF = errors.New("EOF")
		if err == io.EOF {
			return count, nil
		}

		if err != nil {
			return 0, err
		}
	}
}

// io.Writer
// - An interface that has one method `Write`
// - Write(p []byte) (n int, err error)
// - It returns the no. of bytes written and an error
// - If written bytes are less than the provided bytes; you should return an error
// - Common implementation
//   - `*os.File`: for writing to files
//   - `os.Stdout`: writing to standard out
//   - `http.ResponseWriter`: writing the response body of HTTP requests
func writeBytes(s string, w io.Writer) (int, error) {
	n, err := w.Write([]byte(s))
	if err != nil {
		return 0, fmt.Errorf("Error while writing: %w", err)
	}

	return n, nil
}
