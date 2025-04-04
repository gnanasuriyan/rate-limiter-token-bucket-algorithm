package main

import (
	"fmt"
	"github.com/gnanasuriyan/rate-limiter-token-bucket-algorithm/ratelimiter"
	"time"
)

func main() {
	// Create a rate limiter that allows 2 requests per second
	// with a maximum burst capacity of 5 tokens
	limiter := ratelimiter.NewTokenBucket(2, 5)

	// Example: Process requests
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Rate limited\n", i+1)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
