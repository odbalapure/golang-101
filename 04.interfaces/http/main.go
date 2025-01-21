package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Failed to ping google!")
		os.Exit(1)
		return
	}

	// We are logging a struct here, see this: https://pkg.go.dev/net/http
	// Some of the values include
	// Status	  	string
	// StatusCode 	int
	// Proto      	string
	// Header     	Header
	// Body			io.ReaderCloser

	// io.ReaderColser has slightly different syntax we have seen so far
	// Reader		io.Reader -> Read([]byte)(int, error)
	// Closer		io.Closer -> Close() (error)

	// Reader has common o/p for all i.e. []byte
	// HTTP request body
	// Text file / image from harddrive
	// Entering text in command line
	// Data from analog sensor plugged into machine
	// fmt.Println(resp)

	// Get an HTML page from response body
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	// This code returns an HTML page as well
	// NOTE: Copying data from "resp" to the *stdout*
	// func Copy(dst Writer, src Reader) (written int64, err error)
	io.Copy(os.Stdout, resp.Body)
}
