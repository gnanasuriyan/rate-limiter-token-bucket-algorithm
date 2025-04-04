package ratelimiter

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	// Create a token bucket with rate of 2 tokens per second and capacity of 5
	tb := NewTokenBucket(2, 5)

	// Test initial state
	if !tb.Allow() {
		t.Error("Initial request should be allowed")
	}

	// Test rate limiting
	allowed := 0
	for i := 0; i < 10; i++ {
		if tb.Allow() {
			allowed++
		}
		time.Sleep(50 * time.Millisecond)
	}

	// We should have allowed approximately 2 requests per second
	// Given we waited 1 second total (10 * 100ms), we should have allowed around 2 requests
	if allowed < 2 || allowed > 5 {
		t.Errorf("Expected around 5 requests to be allowed, got %d", allowed)
	}

	// Test bucket refill
	time.Sleep(1 * time.Second)

	// After waiting 1 second, we should be able to make at least 1 request
	if !tb.Allow() {
		t.Error("Request should be allowed after refill")
	}
}
