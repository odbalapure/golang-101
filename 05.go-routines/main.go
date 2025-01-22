package main

// Concurrency: When multiple threads executing code. If one thread blocks, another one is picked up & worked upon.
// Parallelism: Multiple threads executed at the same time. This requires multiple CPUs.

import (
	"fmt"
	"net/http"
	"time"
)

// There are 2 types of routine:
// - main: created when we launch our program
// - child: created using the "go" keyword
//
// *Syntax to create a go-routine*
// => go funcName()
// NOTE: Can only be used infront of function calls
//
// Go scheduler runs one thread on each logical core
// By default Go tries to use one core
//
// The *main* routine checks if there is any code it can run
// If not then it exits and the program terminates
// We use channels, so that the *main routine* is aware of - when each of the *child routines* have completed their code.
// We send data to a channel, any running routine can access that data that has access to that channe.
// A channel can be of single type only, ex: a string channel can only send string messages through this channel to go routines.
// It is sort of a two way messaging device.
//
// *Creating a channel*
// => c := make(chan <data-type>)
//
// Sending values in to the channel
// Eg: channel <- 5
// Wait for a value to be sent into the channel. When we get one, assign it to myNumber.
// Eg: myNumber <- channel
// Wait for a value to be sent into the channel. When we get one, log it out.
// Eg: fmt.Println(<- channel)
// NOTE: *channel* stands for channel name

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
		"http://crashed-website.com",
	}

	// Pass the channel to the "checkLink" function
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// NOTE: Not using a loop will result in processing the value from "c" channel only once.
	//		Hence only one value will be printed to the console that gets processed first.
	// fmt.Println(<-c)

	// NOTE: Recieving a value from a channel is a blocking code
	// This seems like a hackish code but IT IS NOT
	// 6 goroutines are created in total (one for each link in the links slice).
	// We loop over the channel to receive the results from these 6 goroutines.
	// The channel blocks until each goroutine sends a value back, so your program will wait until all 6 results are received before printing them.
	// for range links {
	// 	fmt.Println(<-c)
	// }

	// Infinitely pining links mentioned in the slice
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alternative syntax
	for link := range c {
		// The intention it to make a routine put to sleep & immediately start the next goroutine to fetch this link.
		// But this makes the "main" routine do something every 5 seconds.
		// *Sleep* could have been placed inside "checkClink", but the function should only serve its intended purpose.
		// time.Sleep(5 * time.Second)
		// go checkLink(link, c)

		// Theres a better way to do this by using "function literal"
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}()

		// NOTE: This is needed coz the go-routine might read string from a wrong address
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// Some other part of the code sends a value
		// This unblocks the goroutine
		fmt.Println("Link is down", link)
		c <- link
		return
	}
	// This unblocks the goroutine
	c <- link
	fmt.Println("Link is up", link)
}
