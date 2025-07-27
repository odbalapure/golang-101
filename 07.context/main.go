package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// ** Using context for timeout
	openConnectionWithTimeout()

	// ** Using contexts to pass data
	ctx := context.WithValue(context.Background(), "token", "xyz.123.abc")
	result := ctx.Value("token")
	if result == nil {
		fmt.Println("Token not found")
	}

	var token string = result.(string)
	fmt.Println("The token is:", token)
}

// Using a middleware function to pass "userName" around
func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("userId")
		// password := r.FormValue("password")

		// Getting value from a database
		user = "Omi"

		ctx := context.WithValue(r.Context(), "userName", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// Open a connection to a streaming service
// This function is unreliable, it might or might not establish a connection
// Its better to use it with a context with a timeout to avoid the function getting hanged for a long time
func openConnection(done chan bool) {
	fmt.Println("Attempting connection...")
	if rand.Intn(100) > 50 {
		// fmt.Println("OOPS, hanging connection!")
		time.Sleep(10 * time.Hour)
	} else {
		time.Sleep(2 * time.Second)
		// fmt.Println("Connection established!")
	}
	done <- true
}

func openConnectionWithTimeout() {
	// This creates an empty context -> context.Background()
	// Creating a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// If we finish our task before the timeout, we terminate the context
	defer cancel()

	fmt.Println(ctx.Deadline())

	done := make(chan bool)
	go openConnection(done)

	// NOTE: select is strictly for channel operations
	select {
	case <-done:
		fmt.Println("Connection successful.")
	case <-ctx.Done():
		fmt.Println("Context connection got timed out :(")
	}
}
