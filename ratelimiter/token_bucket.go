package ratelimiter

import (
	"sync"
	"time"
)

// TokenBucket represents a token bucket rate limiter
type TokenBucket struct {
	rate       float64    // tokens per second
	capacity   float64    // maximum number of tokens
	tokens     float64    // current number of tokens
	lastRefill time.Time  // last time tokens were refilled
	mu         sync.Mutex // mutex for thread safety
}

// NewTokenBucket creates a new token bucket rate limiter
func NewTokenBucket(rate float64, capacity float64) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity, // start with a full bucket
		lastRefill: time.Now(),
	}
}

// refill adds tokens to the bucket based on the time elapsed since the last refill
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	// Calculate new tokens to add
	newTokens := elapsed * tb.rate
	// Update tokens and last refill time
	tb.tokens = min(tb.capacity, tb.tokens+newTokens)
	tb.lastRefill = now
}

// Allow checks if a request is allowed and consumes a token if it is
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Refill the bucket
	tb.refill()

	// Check if we have enough tokens
	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
		return true
	}

	return false
}

// min returns the minimum of two float64 values
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
