## Context

### Using conext with a timeout

```go
func main() {
	openConnectionWithTimeout()
}

// Open a connection to a streaming service
// This function is unreliable, it might or might not establish a connection
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
```

**Context Deadline**

> `ctx.Deadline()` shows when the conext will expire

**Context Error**
> `ctx.Err()` can be used to get the context error directly

### Using conext to share data

