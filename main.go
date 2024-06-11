package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	capacity        int
	numOfTokens     int
	refillFrequency int
	lastRefillStamp time.Time
}

func newTokenBucket(capacity, refillFrequency int) *TokenBucket {
	return &TokenBucket{
		capacity:        capacity,
		numOfTokens:     capacity,
		refillFrequency: refillFrequency,
		lastRefillStamp: time.Now(),
	}
}

func (bucket *TokenBucket) canProcessRequests(numOfRequests int) bool {
	fmt.Println(bucket.capacity)
	fmt.Println(bucket.lastRefillStamp)

	if numOfRequests <= bucket.numOfTokens {
		bucket.numOfTokens -= numOfRequests
		return true
	}

	return false

}

func tokenBucket() {
	bucket := newTokenBucket(10, 3)
	fmt.Println("token bucket algorithm")
	bucket.canProcessRequests(500)

}

func main() {

	fmt.Println("rate limit algorithms")
	tokenBucket()

}
